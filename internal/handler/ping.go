package handler

import (
	"context"
	"vague/internal/process"
	"vague/internal/runtime"
)

func init() {
	runtime.Register("ping", runtime.HandlerMeta{
		Desc:    "Ping the runtime",
		Args:    nil,
		Fn:      Ping,
	})
}

type PingResponse struct {
	Message string `json:"message"`
}

// Ping goes through the runtime so a reply proves the runtime loop is alive,
// not just the socket.
func Ping(ctx context.Context, sess runtime.Session, params process.ExecuteParams, rt *runtime.Runtime) (any, error) {
	return rt.Do(ctx, func(st *runtime.State) (any, error) {
		return &PingResponse{Message: "pong"}, nil
	})
}
