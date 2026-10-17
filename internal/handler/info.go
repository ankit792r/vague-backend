package handler

import (
	"context"
	"vague/internal/process"
	"vague/internal/runtime"
)

func init() {
	runtime.Register("info", SystemInfo)
}

func SystemInfo(ctx context.Context, sess runtime.Session, params process.ExecuteParams, rt *runtime.Runtime) (any, error) {
	return rt.Do(ctx, func(st *runtime.State) (any, error) {
		return "SystemInfo OK", nil
	})
}
