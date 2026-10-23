package buffer

import (
	"vague/internal/runtime"
)

const CMD_ScratchBuffer = "scratch-buffer"

var scratchBuffer = runtime.Command{
	Name: CMD_ScratchBuffer,
	Desc: "New scratch buffer",
	Args: runtime.NoArgs,
	RunC: func(env runtime.ExeEnv, args []string) (any, error) {
		sess := env.Sess

		return env.Rt.Do(env.Ctx, func(st *runtime.State) (any, error) {
			buff := NewScratch(sess.NextBufferId(), "scratch")
		
			
		})
	},
}
