package commands

import (
	"fmt"
	"vague/internal/runtime"
)

func init() {
	runtime.Register(cmd_getBuiltinCommand)
	runtime.Register(cmd_listBuiltinCommands)
}

const (
	CMD_ListCommands = "list-commands"
	CMD_GetCommand   = "get-command"
)

var cmd_listBuiltinCommands = runtime.Command{
	Name: CMD_ListCommands,
	Desc: "get all commands name",
	Args: runtime.NoArgs,
	RunC: func(env runtime.ExeEnv, args []string) (any, error) {
		return env.Rt.Do(env.Ctx, func(_ *runtime.State) (any, error) {
			return runtime.AllCommandsNames(), nil
		})
	},
}

type getbuiltincmd_res struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

var cmd_getBuiltinCommand = runtime.Command{
	Name: CMD_GetCommand,
	Desc: "get command info",
	Args: runtime.ExactArgs(1),
	RunC: func(env runtime.ExeEnv, args []string) (any, error) {
		return env.Rt.Do(env.Ctx, func(st *runtime.State) (any, error) {
			cmd, ok := runtime.Lookup(args[0])
			if !ok {
				return nil, fmt.Errorf("requested command does not exists in lookup table")
			}

			res := getbuiltincmd_res{
				Name: cmd.Name,
				Desc: cmd.Desc,
			}

			return res, nil
		})
	},
}
