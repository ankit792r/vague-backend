package runtime

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"sync"
	"vague/features/buffer"
	"vague/internal/process"
)

type Session interface {
	ID() uint64
	Send(msg process.Message) bool
	NextWindowId() uint64
	NextBufferId() uint64

	Buffers map[uint64]*buffer.Buffer
}

type ExeEnv struct {
	Ctx  context.Context
	Sess Session
	Rt   *Runtime
}

type PositionalArgs func(args []string) error
type CommandFn func(env ExeEnv, args []string) (any, error)

var NoArgs PositionalArgs = func(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("takes no arguments, received %d", len(args))
	}
	return nil
}

var ArbitraryArgs PositionalArgs = func(args []string) error {
	return nil
}

func MinimumNArgs(count int) PositionalArgs {
	return func(args []string) error {
		if len(args) < count {
			return fmt.Errorf(
				"accepts at least %d arg(s), received %d",
				count,
				len(args),
			)
		}
		return nil
	}
}

func MaximumNArgs(count int) PositionalArgs {
	return func(args []string) error {
		if len(args) > count {
			return fmt.Errorf(
				"accepts %d arg(s) max, received %d",
				count,
				len(args),
			)
		}
		return nil
	}
}

func ExactArgs(count int) PositionalArgs {
	return func(args []string) error {
		if len(args) != count {
			return fmt.Errorf(
				"accepts %d arg(s), received %d",
				count,
				len(args),
			)
		}

		return nil
	}
}

type Command struct {
	Name string
	Desc string
	Args PositionalArgs
	RunC CommandFn
}

var (
	commandMu sync.RWMutex
	commands  = make(map[string]Command)
)

func (c *Command) Execute(env ExeEnv, args []string) (any, error) {
	if err := c.Args(args); err != nil {
		return nil, err
	}
	return c.RunC(env, args)
}

func AllCommandsNames() []string {
	return slices.Collect(maps.Keys(commands))
}

// --------------------------------------------

var pingCmd = Command{
	Name: "ping",
	Desc: "Send ping command to runtime",
	Args: NoArgs,
	RunC: func(env ExeEnv, args []string) (any, error) {
		return "Pong", nil
	},
}

func Register(cmd Command) {
	if cmd.Name == "" {
		panic("runtime: Register with empty method name")
	}
	if cmd.RunC == nil {
		panic(fmt.Sprintf("runtime: Register %q with nil handler", cmd.Name))
	}

	commandMu.Lock()
	defer commandMu.Unlock()

	if _, dup := commands[cmd.Name]; dup {
		panic(fmt.Sprintf("runtime: Register called twice for method %q", cmd.Name))
	}

	commands[cmd.Name] = cmd
}

func Lookup(method string) (Command, bool) {
	commandMu.RLock()
	defer commandMu.RUnlock()

	cmd, ok := commands[method]
	return cmd, ok
}

func init() {
	Register(pingCmd)
}
