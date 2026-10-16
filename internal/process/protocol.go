package process

const (
	// MethodExecute runs a registered editor command. Request.
	MethodExecute = "execute"

	// MethodInput delivers keys. Notification, client to server.
	MethodInput = "input"
)

// ExecuteParams invokes a registered editor command.
type ExecuteParams struct {
	Name  string   `json:"name"`
	Args  []string `json:"args,omitempty"`
	Bang  bool     `json:"bang,omitempty"`
	Count int      `json:"count,omitempty"`
}
