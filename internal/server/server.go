package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"vague/internal/process"
	"vague/internal/runtime"

	_ "vague/internal/handler"
)

type Server struct {
	wg sync.WaitGroup
	mu sync.Mutex

	sessions map[*Session]struct{}
	runtime  *runtime.Runtime

	nextSessionID uint64
}

func New() *Server {
	return &Server{
		sessions: make(map[*Session]struct{}),
		runtime:  runtime.New(),
	}
}

func (s *Server) Serve(ctx context.Context, listener *process.Listener) error {
	defer listener.Close()

	go s.runtime.Run()

	go func() {
		<-ctx.Done()
		listener.CloseV2()

	}()

	accError := s.acceptConnections(ctx, listener)

	s.wg.Wait()

	return accError
}

func (s *Server) acceptConnections(cxt context.Context, listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) || cxt.Err() != nil {
				return nil
			}

			return fmt.Errorf("accept: %w", err)
		}

		s.nextSessionID++
		sess := NewSession(s.nextSessionID, conn)

		s.addSession(sess)

		s.wg.Go(func() {
			defer s.removeSession(sess)
			s.ServeSess(cxt, sess)
		})
	}
}

// Session bookkeeping
func (s *Server) addSession(sess *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[sess] = struct{}{}
}

func (s *Server) removeSession(sess *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, sess)
}

func (s *Server) closeAllSessions() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for sess := range s.sessions {
		sess.Conn.Close()
	}
}
