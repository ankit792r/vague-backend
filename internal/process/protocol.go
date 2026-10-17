package process

const (
	ActionExecute = "execute"
	ActionInput   = "input"
	ActionNotify  = "notify"
)

// ExecuteParams invokes a registered editor command.
type ExecuteParams struct {
	Name  string   `json:"name"`
	Args  []string `json:"args,omitempty"`
	Bang  bool     `json:"bang,omitempty"`
	Count int      `json:"count,omitempty"`
}
