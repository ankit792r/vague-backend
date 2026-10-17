package runtime

import (
	"context"
	"fmt"
)

func init() {
	Register("function_list", "List all functions", function_names)
	Register("function_count", "Get the number of functions", function_count)
	Register("function_exist", "Check if a function exists", function_exist)
	Register("function_describe", "Describe a function and its arguments", function_describe)
}

type fnamesResponse struct {
	Names []string `json:"names"`
}

func function_names(ctx context.Context, sess Session, req Request[NoArgs], rt *Runtime) (any, error) {
	all := Handlers()

	names := make([]string, len(all))
	for i, h := range all {
		names[i] = h.Name
	}

	return &fnamesResponse{Names: names}, nil
}

type fcountResponse struct {
	Count int `json:"count"`
}

func function_count(ctx context.Context, sess Session, req Request[NoArgs], rt *Runtime) (any, error) {
	return &fcountResponse{Count: len(Handlers())}, nil
}

type fexistsResponse struct {
	Exist bool `json:"exist"`
}

// FunctionNameArgs names the function a helper operates on.
type FunctionNameArgs struct {
	Name string `json:"name" desc:"Name of the function" required:"true"`
}

func function_exist(ctx context.Context, sess Session, req Request[FunctionNameArgs], rt *Runtime) (any, error) {
	_, exist := Lookup(req.Args.Name)
	return &fexistsResponse{Exist: exist}, nil
}

type fdescribeResponse struct {
	Description string    `json:"description"`
	Args        []ArgMeta `json:"args"`
}

func function_describe(ctx context.Context, sess Session, req Request[FunctionNameArgs], rt *Runtime) (any, error) {
	hmeta, ok := Describe(req.Args.Name)
	if !ok {
		return nil, fmt.Errorf("unknown function %q", req.Args.Name)
	}

	return &fdescribeResponse{Description: hmeta.Desc, Args: hmeta.Args}, nil
}
