package runtime

import (
	"context"
)

type ExeEnv struct {
	Ctx  context.Context
	Sess Session
	Rt   *Runtime
}

type CommandFn func(env ExeEnv, args []string) (any, error)

type Command struct {
	Name string
	Desc string
	Args []string
	RunC CommandFn
}

var Commands = make(map[string]Command)

func (c *Command) Execute(env ExeEnv, args []string) (any, error) {
	// Here we will do some args parsing
	// such as NoArgs or ExactArgs
	return c.RunC(env, args)
}

// --------------------------------------------

var pingCmd = Command{
	Name: "ping",
	Desc: "Send ping command to runtime",
	Args: []string{},
	RunC: func(env ExeEnv, args []string) (any, error) {
		return "Pong", nil
	},
}

func init() {
	Commands[pingCmd.Name] = pingCmd
}
