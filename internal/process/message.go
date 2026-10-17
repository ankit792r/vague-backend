package process

import (
	"encoding/json"
	"fmt"
)

type Kind string

const (
	// KindRequest expects exactly one KindResponse with a matching ID.
	KindRequest Kind = "request"

	// KindResponse answers a request.
	KindResponse Kind = "response"
)

type Message struct {
	Kind Kind `json:"kind"`

	// ID correlates a request with its response.
	ID uint64 `json:"id"`

	Action string `json:"action,omitempty"`

	Params json.RawMessage `json:"params,omitempty"`

	Result json.RawMessage `json:"result,omitempty"`
	Error  string          `json:"error,omitempty"`
}

// NewRequest builds a request, encoding its parameters.
func NewRequest(id uint64, action string, params any) (Message, error) {
	raw, err := Encode(params)
	if err != nil {
		return Message{}, err
	}

	return Message{Kind: KindRequest, ID: id, Action: action, Params: raw}, nil
}

// NewResponse builds a successful response.
func NewResponse(id uint64, result any) (Message, error) {
	raw, err := Encode(result)
	if err != nil {
		return Message{}, err
	}

	return Message{Kind: KindResponse, ID: id, Result: raw}, nil
}

// NewErrorResponse builds a failed response.
func NewErrorResponse(id uint64, cause error) Message {
	return Message{Kind: KindResponse, ID: id, Error: cause.Error()}
}

// Validate checks that a decoded message is internally consistent, so that
// handlers can assume the fields they care about are present.
func (m Message) Validate() error {
	switch m.Kind {
	case KindRequest:
		if m.ID == 0 {
			return fmt.Errorf("request %q has no id", m.Action)
		}
		if m.Action == "" {
			return fmt.Errorf("request %d has no action", m.ID)
		}

	case KindResponse:
		if m.ID == 0 {
			return fmt.Errorf("response has no id")
		}

	default:
		return fmt.Errorf("unknown message kind %q", m.Kind)
	}

	return nil
}

// Encode marshals a value for a Params or Result field. Payloads are carried
// as raw JSON so that neither side has to marshal a value only to unmarshal
// it again into the type it wanted.
func Encode(v any) (json.RawMessage, error) {
	if v == nil {
		return nil, nil
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("encode payload: %w", err)
	}

	return raw, nil
}

// DecodeParams unmarshals a request or notification's parameters into dst.
// Absent parameters leave dst untouched, so methods that take none can skip
// the call.
func (m Message) DecodeParams(dst any) error {
	if len(m.Params) == 0 {
		return nil
	}

	if err := json.Unmarshal(m.Params, dst); err != nil {
		return fmt.Errorf("decode params for %q: %w", m.Action, err)
	}

	return nil
}

// DecodeResult unmarshals a response's result into dst.
func (m Message) DecodeResult(dst any) error {
	if len(m.Result) == 0 {
		return nil
	}

	if err := json.Unmarshal(m.Result, dst); err != nil {
		return fmt.Errorf("decode result: %w", err)
	}

	return nil
}
