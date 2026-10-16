package server

import (
	"context"
	"fmt"
	"vague/internal/process"
)

func (s *Server) dispatch(_ context.Context, _ *Session, msg process.Message) {
	fmt.Println(msg)
}
