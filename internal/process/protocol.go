package process

import "encoding/json"

const (
	ActionExecute = "execute"
	ActionInput   = "input"
	ActionNotify  = "notify"
)

// ExecuteParams invokes a registered editor command. Args is kept raw so the
// handler can decode it into its own argument type.
type ExecuteParams struct {
	Name  string          `json:"name"`
	Args  json.RawMessage `json:"args,omitempty"`
	Bang  bool            `json:"bang,omitempty"`
	Count int             `json:"count,omitempty"`
}

// NewExecuteParams builds the parameters for invoking name with args, which
// may be nil for commands that take none.
func NewExecuteParams(name string, args any) (ExecuteParams, error) {
	raw, err := Encode(args)
	if err != nil {
		return ExecuteParams{}, err
	}

	return ExecuteParams{Name: name, Args: raw}, nil
}
