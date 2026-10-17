package client

import (
	"context"
	"errors"
	"fmt"
	"vague/internal/process"
)

func (c *Client) call(ctx context.Context, action string, params any, result any) error {
	msg, err := func() (process.Message, error) {

		c.mu.Lock()
		defer c.mu.Unlock()

		// check error
		if c.err != nil {
			return process.Message{}, c.err
		}

		id := c.nextId
		c.nextId++

		msg, err := process.NewRequest(id, action, params)
		if err != nil {
			return process.Message{}, err
		}

		c.pending[id] = make(chan process.Message, 1)

		return msg, nil
	}()
	if err != nil {
		return err
	}

	waiter := c.waiterFor(msg.ID)

	if err := c.write(msg); err != nil {
		c.forget(msg.ID)
		return fmt.Errorf("send %s: %w", action, err)
	}

	select {
	case res, ok := <-waiter:
		if !ok {
			return c.reason()
		}

		if res.Error != "" {
			return errors.New(res.Error)
		}

		if result == nil {
			return nil
		}

		return res.DecodeResult(result)

	case <-ctx.Done():
		c.forget(msg.ID)
		return ctx.Err()

	case <-c.done:
		c.forget(msg.ID)
		return ctx.Err()
	}
}

// Call sends a request for method and decodes its result into result, which
// may be nil to discard it.
func (c *Client) Call(ctx context.Context, method string, params any, result any) error {
	return c.call(ctx, method, params, result)
}

// Execute runs a named editor command and decodes its result into result,
// which may be nil to discard it.
func (c *Client) Execute(ctx context.Context, params process.ExecuteParams, result any) error {
	return c.call(ctx, process.ActionExecute, params, result)
}
