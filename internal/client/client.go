package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"vague/internal/process"
)

var ErrClosed = errors.New("connection closed")

type Client struct {
	conn net.Conn

	reader  *bufio.Reader
	pending map[uint64]chan process.Message

	mu      sync.Mutex
	writeMu sync.Mutex
	nextId  uint64

	err error

	done      chan struct{}
	closeOnce sync.Once
}

func Connect() (*Client, error) {
	conn, err := process.Dial()
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:    conn,
		done:    make(chan struct{}),
		reader:  bufio.NewReader(conn),
		pending: make(map[uint64]chan process.Message),
		nextId:  1,
	}

	go client.readLoop()

	return client, nil
}

func (c *Client) Close() error {
	c.closeOnce.Do(func() {
		close(c.done)
	})

	return c.conn.Close()
}

func (c *Client) readLoop() {
	for {
		payload, err := process.ReadFrame(c.reader)
		if err != nil {
			c.fail(err)
			return
		}

		var msg process.Message
		if err := json.Unmarshal(payload, &msg); err != nil {
			c.fail(fmt.Errorf("decode message: %w", err))
			return
		}

		switch msg.Kind {
		case process.KindResponse:
			c.deliver(msg)

		default:
			// The server does not call the client, so this is a protocol
			// violation rather than something to route.
			c.fail(fmt.Errorf("unexpected message kind %q from server", msg.Kind))
			return
		}

	}
}

// deliver hands a response to the call waiting on its ID.
func (c *Client) deliver(msg process.Message) {
	c.mu.Lock()
	waiter, ok := c.pending[msg.ID]
	delete(c.pending, msg.ID)
	c.mu.Unlock()

	if !ok {
		// A call that gave up, or a reply to a malformed frame, which
		// carries no usable ID.
		return
	}

	waiter <- msg
}

func (c *Client) write(msg process.Message) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()

	return process.WriteFrame(c.conn, msg)
}

func (c *Client) waiterFor(id uint64) chan process.Message {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.pending[id]
}

func (c *Client) forget(id uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.pending, id)
}

func (c *Client) reason() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err != nil {
		return c.err
	}

	return ErrClosed
}

// fail records why the connection ended and wakes every waiting call.
func (c *Client) fail(cause error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err == nil {
		if errors.Is(cause, io.EOF) || errors.Is(cause, net.ErrClosed) {
			c.err = ErrClosed
		} else {
			c.err = cause
		}
	}

	for id, waiter := range c.pending {
		close(waiter)
		delete(c.pending, id)
	}
}
