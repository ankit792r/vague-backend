package handler

import (
	"context"
	"vague/internal/runtime"
)

func init() {
	runtime.Register("ping", "Ping the runtime", Ping)
}

type PingResponse struct {
	Message string `json:"message"`
}

// Ping goes through the runtime so a reply proves the runtime loop is alive,
// not just the socket.
func Ping(ctx context.Context, sess runtime.Session, req runtime.Request[runtime.NoArgs], rt *runtime.Runtime) (any, error) {
	return rt.Do(ctx, func(st *runtime.State) (any, error) {
		return &PingResponse{Message: "pong"}, nil
	})
}
