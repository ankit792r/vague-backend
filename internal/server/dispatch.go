package server

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"vague/internal/process"
	"vague/internal/runtime"
)

func (s *Server) dispatch(ctx context.Context, sess *Session, msg process.Message) {
	if msg.Kind != process.KindRequest {
		slog.Debug("ignoring non-request message", "kind", msg.Kind, "id", msg.ID)
		return
	}

	handle, ok := runtime.Lookup(msg.Method)
	if !ok {
		sess.Reply(msg.ID, nil, fmt.Errorf("unknown method %q", msg.Method))
		return
	}

	result, err := s.call(ctx, handle, sess, msg)
	sess.Reply(msg.ID, result, err)
}

// call runs a handler, turning a panic outside rt.Do into an error reply
// instead of taking the whole server down.
func (s *Server) call(ctx context.Context, handle runtime.HandlerFunc, sess *Session, msg process.Message) (result any, err error) {
	defer func() {
		if p := recover(); p != nil {
			slog.Error("handler panicked",
				"method", msg.Method,
				"panic", p,
				"stack", string(debug.Stack()),
			)
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	return handle(ctx, sess, msg, s.runtime)
}
