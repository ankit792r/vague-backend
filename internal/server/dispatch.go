package server

import (
	"context"
	"fmt"
	"log/slog"
	"vague/internal/process"
	"vague/internal/runtime"
)

func (s *Server) dispatch(ctx context.Context, sess *Session, msg process.Message) {
	if msg.Kind != process.KindRequest {
		slog.Debug("ignoring non-request message", "kind", msg.Kind, "id", msg.ID)
		return
	}

	var param process.CommandParam
	var cmd runtime.Command

	if msg.Method == process.MethodExecute {
		if err := msg.DecodeParams(&param); err != nil {
			sess.Reply(msg.ID, nil, fmt.Errorf("faild to parse args for method %q", msg.Method))
		}
		cmdF, ok := runtime.Lookup(param.Name)
		if !ok {
			sess.Reply(msg.ID, nil, fmt.Errorf("unknown command %q", param.Name))
			return
		}
		cmd = cmdF
	}

	exEnv := runtime.ExeEnv{
		Ctx:  ctx,
		Sess: sess,
		Rt:   s.runtime,
	}

	result, err := cmd.Execute(exEnv, param.Args)
	sess.Reply(msg.ID, result, err)
}
