package runtime

import (
	"context"
	"fmt"
	"sync"
	"vague/internal/process"
)

// Session is the view of a client connection that handlers get. It is an
// interface so that handlers do not depend on the server package, which
// imports them.
type Session interface {
	ID() uint64
	Send(msg process.Message) bool
}

// HandlerFunc serves one request method. Its result becomes the response
// payload; editor state must only be touched inside rt.Do.
type HandlerFunc func(ctx context.Context, sess Session, params process.ExecuteParams, rt *Runtime) (any, error)

var (
	handlersMu sync.RWMutex
	handlers   = make(map[string]HandlerFunc)
)

// Register binds a handler to a method name. It is meant to be called from
// init, and panics on misuse so a bad registration fails at startup.
func Register(method string, fn HandlerFunc) {
	if method == "" {
		panic("runtime: Register with empty method name")
	}
	if fn == nil {
		panic(fmt.Sprintf("runtime: Register %q with nil handler", method))
	}

	handlersMu.Lock()
	defer handlersMu.Unlock()

	if _, dup := handlers[method]; dup {
		panic(fmt.Sprintf("runtime: Register called twice for method %q", method))
	}

	handlers[method] = fn
}

// Lookup returns the handler registered for method.
func Lookup(method string) (HandlerFunc, bool) {
	handlersMu.RLock()
	defer handlersMu.RUnlock()

	fn, ok := handlers[method]
	return fn, ok
}
