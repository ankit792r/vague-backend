package runtime

import (
	"context"
	"maps"
	"slices"
	"vague/internal/process"
)

func init() {
	Register("function_list", function_names)
	Register("function_count", function_count)
	Register("function_exist", function_exist)
	// Register("function_desc", function_desc)
	// Register("function_args", function_args)
}

type fnamesResponse struct {
	Names []string `json:"names"`
}

func function_names(ctx context.Context, sess Session, params process.ExecuteParams, rt *Runtime) (any, error) {
	names := slices.Collect(maps.Keys(handlers))
	return &fnamesResponse{Names: names}, nil
}

type fcountResponse struct {
	Count int `json:"count"`
}

func function_count(ctx context.Context, sess Session, params process.ExecuteParams, rt *Runtime) (any, error) {
	return &fcountResponse{Count: len(handlers)}, nil
}

type fexistsResponse struct {
	Exist bool `json:"exist"`
}

func function_exist(ctx context.Context, sess Session, params process.ExecuteParams, rt *Runtime) (any, error) {
	_, exist := Lookup(params.Args[0])
	return &fexistsResponse{Exist: exist}, nil
}

// type fdescResponse struct {
// 	Description string `json:"description"`
// }

// func function_desc(ctx context.Context, sess Session, params process.ExecuteParams, rt *Runtime) (any, error) {
// 	desc, exist := Lookup(params.Args[0])
// 	return &fdescResponse{Description: desc.Description}, nil
// }

// type fargsResponse struct {
// 	Args []string `json:"args"`
// }

// func function_args(ctx context.Context, sess Session, params process.ExecuteParams, rt *Runtime) (any, error) {
// 	args, exist := Lookup(params.Args[0])
// 	return &fargsResponse{Args: args.Args}, nil
// }
