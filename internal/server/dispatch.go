package server

import (
	"context"
	"fmt"
	"vague/internal/process"
)

func (s *Server) dispatch(_ context.Context, sess *Session, msg process.Message) {
	fmt.Println(msg)
	sess.Reply(msg.ID, "pong", nil)
}
