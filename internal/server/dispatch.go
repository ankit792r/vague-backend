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

	var params process.ExecuteParams
	var handle runtime.HandlerFunc
	var ok bool

	if msg.Action == process.ActionExecute {
		if err := msg.DecodeParams(&params); err != nil {
			sess.Reply(msg.ID, nil, err)
			return
		}

		handle, ok = runtime.Lookup(params.Name)
		if !ok {
			sess.Reply(msg.ID, nil, fmt.Errorf("unknown method %q", params.Name))
			return
		}
	}

	result, err := s.call(ctx, handle, sess, params)
	sess.Reply(msg.ID, result, err)
}

// call runs a handler, turning a panic outside rt.Do into an error reply
// instead of taking the whole server down.
func (s *Server) call(ctx context.Context, handle runtime.HandlerFunc, sess *Session, params process.ExecuteParams) (result any, err error) {
	defer func() {
		if p := recover(); p != nil {
			slog.Error("handler panicked",
				"method", params.Name,
				"panic", p,
				"stack", string(debug.Stack()),
			)
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	return handle(ctx, sess, params, s.runtime)
}
