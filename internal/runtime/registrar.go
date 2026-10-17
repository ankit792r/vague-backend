package runtime

import (
	"cmp"
	"context"
	"fmt"
	"reflect"
	"slices"
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

// Request is what a typed handler receives, with Args already decoded.
type Request[A any] struct {
	Name  string
	Args  A
	Bang  bool
	Count int
}

// Handler is a handler whose arguments are decoded into A, which must be a
// struct. Use NoArgs for handlers that take none.
type Handler[A any] func(ctx context.Context, sess Session, req Request[A], rt *Runtime) (any, error)

type HandlerMeta struct {
	Name string      `json:"name"`
	Desc string      `json:"description"`
	Args []ArgMeta   `json:"args,omitempty"`
	Fn   HandlerFunc `json:"-"`
}

var (
	handlersMu sync.RWMutex
	handlers   = make(map[string]HandlerMeta)
)

// Register binds a handler to a method name, deriving its argument metadata
// from A. It is meant to be called from init, and panics on misuse so a bad
// registration fails at startup.
func Register[A any](method, desc string, fn Handler[A]) {
	if method == "" {
		panic("runtime: Register with empty method name")
	}
	if fn == nil {
		panic(fmt.Sprintf("runtime: Register %q with nil handler", method))
	}

	t := reflect.TypeFor[A]()
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("runtime: Register %q: args type %s is not a struct", method, t))
	}

	args := describeArgs(t)

	var required []string
	for _, a := range args {
		if a.Required {
			required = append(required, a.Name)
		}
	}

	hmeta := HandlerMeta{
		Name: method,
		Desc: desc,
		Args: args,
		Fn: func(ctx context.Context, sess Session, params process.ExecuteParams, rt *Runtime) (any, error) {
			decoded, err := decodeArgs[A](params.Args, required)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", method, err)
			}

			return fn(ctx, sess, Request[A]{
				Name:  params.Name,
				Args:  decoded,
				Bang:  params.Bang,
				Count: params.Count,
			}, rt)
		},
	}

	handlersMu.Lock()
	defer handlersMu.Unlock()

	if _, dup := handlers[method]; dup {
		panic(fmt.Sprintf("runtime: Register called twice for method %q", method))
	}

	handlers[method] = hmeta
}

// Lookup returns the handler registered for method.
func Lookup(name string) (HandlerFunc, bool) {
	hmeta, ok := Describe(name)
	return hmeta.Fn, ok
}

// Describe returns the metadata registered for method.
func Describe(name string) (HandlerMeta, bool) {
	handlersMu.RLock()
	defer handlersMu.RUnlock()

	hmeta, ok := handlers[name]
	return hmeta, ok
}

// Handlers returns the metadata of every registered handler, sorted by name.
func Handlers() []HandlerMeta {
	handlersMu.RLock()
	defer handlersMu.RUnlock()

	out := make([]HandlerMeta, 0, len(handlers))
	for _, h := range handlers {
		out = append(out, h)
	}

	slices.SortFunc(out, func(a, b HandlerMeta) int { return cmp.Compare(a.Name, b.Name) })
	return out
}
