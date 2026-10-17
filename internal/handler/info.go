package handler

import (
	"context"
	"vague/internal/process"
	"vague/internal/runtime"
)

func init() {
	runtime.Register("info", runtime.HandlerMeta{
		Desc:    "Get system information",
		Args:    nil,
		Fn:      SystemInfo,
	})
}

type SystemInfoResponse struct {
	Version string `json:"version"`
	Build   string `json:"build"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func SystemInfo(ctx context.Context, sess runtime.Session, params process.ExecuteParams, rt *runtime.Runtime) (any, error) {
	return rt.Do(ctx, func(st *runtime.State) (any, error) {
		return "SystemInfo OK", nil
	})
}
