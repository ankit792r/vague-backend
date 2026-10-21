package server

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"sync"
	"time"
	"vague/internal/process"
)

const outboundQueue = 256

type Session struct {
	Id     uint64
	Conn   net.Conn
	Reader *bufio.Reader

	outbound chan process.Message

	done     chan struct{}
	doneOnce sync.Once
}

func NewSession(id uint64, conn net.Conn) *Session {
	return &Session{
		Id:       id,
		Conn:     conn,
		Reader:   bufio.NewReader(conn),
		outbound: make(chan process.Message, outboundQueue),

		done: make(chan struct{}),
	}
}

func (s *Session) ID() uint64 { return s.Id }

func (s *Session) Close() {
	s.doneOnce.Do(func() {
		close(s.done)
	})
}

func (s *Server) ServeSess(ctx context.Context, sess *Session) {
	ctx, stop := context.WithCancel(ctx)
	defer stop()

	sessionWriteComplete := make(chan struct{})
	go func() {
		defer close(sessionWriteComplete)
		sess.WriteLoop()
	}()

	defer func() {
		sess.Close()

		// Let the writer flush before the connection goes away, but do not
		// hang on a client that has stopped reading.
		select {
		case <-sessionWriteComplete:
		case <-time.After(time.Second):
		}
	}()

	for {

		payload, err := process.ReadFrame(sess.Reader)
		if err != nil {
			if !errors.Is(err, io.EOF) && !errors.Is(err, net.ErrClosed) {
				slog.Debug("session read ended", "error", err)
			}
			break
		}

		var msg process.Message
		if err := json.Unmarshal(payload, &msg); err != nil {
			// The ID is unknown, so the client cannot correlate this. It is
			// still worth sending so a malformed call is not a silent hang.
			sess.Reply(0, nil, fmt.Errorf("malformed message: %w", err))
			continue
		}

		if err := msg.Validate(); err != nil {
			sess.Reply(msg.ID, nil, err)
			continue
		}

		s.dispatch(ctx, sess, msg)
	}
}

func (s *Session) WriteLoop() {
	for {
		select {
		case msg := <-s.outbound:
			if !s.write(msg) {
				return
			}

		case <-s.done:
			// Flush what is already queued so a reply sent just before
			// shutdown still reaches the client.
			for {
				select {
				case msg := <-s.outbound:
					if !s.write(msg) {
						return
					}
				default:
					return
				}
			}
		}
	}
}

func (s *Session) write(msg process.Message) bool {
	if err := process.WriteFrame(s.Conn, msg); err != nil {
		if !errors.Is(err, net.ErrClosed) {
			slog.Debug("write failed", "error", err)
		}

		return false
	}

	return true
}

// send queues a message, reporting whether the session can still keep up.
func (s *Session) Send(msg process.Message) bool {
	select {
	case <-s.done:
		return false
	default:
	}

	select {
	case s.outbound <- msg:
		return true

	case <-s.done:
		return false

	default:
		slog.Warn("outbound queue full, dropping session")
		s.Conn.Close()

		return false
	}
}

// reply answers a request.
func (s *Session) Reply(id uint64, result any, cause error) {
	if cause != nil {
		s.Send(process.NewErrorResponse(id, cause))
		return
	}

	msg, err := process.NewResponse(id, result)
	if err != nil {
		s.Send(process.NewErrorResponse(id, err))
		return
	}

	s.Send(msg)
}
