package runtime

import (
	"context"
	"fmt"
)

func init() {
	Register("function_list", "List the names of all registered functions, sorted alphabetically", function_list)
	Register("function_count", "Get the number of registered functions", function_count)
	Register("function_exist", "Check whether a function with the given name is registered", function_exist)
	Register("function_describe", "Get the name, description and arguments of a registered function", function_describe)
}

// FunctionNameArgs names the function a builtin operates on.
type FunctionNameArgs struct {
	Name string `json:"name" desc:"Name of the registered function" required:"true"`
}

type flistResponse struct {
	Names []string `json:"names"`
}

func function_list(ctx context.Context, sess Session, req Request[NoArgs], rt *Runtime) (any, error) {
	all := Handlers()

	names := make([]string, len(all))
	for i, h := range all {
		names[i] = h.Name
	}

	return &flistResponse{Names: names}, nil
}

type fcountResponse struct {
	Count int `json:"count"`
}

func function_count(ctx context.Context, sess Session, req Request[NoArgs], rt *Runtime) (any, error) {
	return &fcountResponse{Count: len(Handlers())}, nil
}

type fexistResponse struct {
	Exist bool `json:"exist"`
}

func function_exist(ctx context.Context, sess Session, req Request[FunctionNameArgs], rt *Runtime) (any, error) {
	_, exist := Lookup(req.Args.Name)
	return &fexistResponse{Exist: exist}, nil
}

func function_describe(ctx context.Context, sess Session, req Request[FunctionNameArgs], rt *Runtime) (any, error) {
	hmeta, ok := Describe(req.Args.Name)
	if !ok {
		return nil, fmt.Errorf("unknown function %q", req.Args.Name)
	}

	return &hmeta, nil
}
