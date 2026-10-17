package runtime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type ArgType string

const (
	ArgString ArgType = "string"
	ArgInt    ArgType = "int"
	ArgFloat  ArgType = "float"
	ArgBool   ArgType = "bool"
	ArgArray  ArgType = "array"
	ArgObject ArgType = "object"
	ArgAny    ArgType = "any"
)

// ArgMeta describes one argument of a handler. It is derived from the
// handler's argument struct:
//
//	type Args struct {
//		Name string `json:"name" desc:"Function to look up" required:"true"`
//	}
type ArgMeta struct {
	Name     string  `json:"name"`
	Type     ArgType `json:"type"`
	Desc     string  `json:"description,omitempty"`
	Required bool    `json:"required,omitempty"`
}

// NoArgs is the argument type for handlers that take no arguments.
type NoArgs struct{}

func describeArgs(t reflect.Type) []ArgMeta {
	var out []ArgMeta

	for i := range t.NumField() {
		f := t.Field(i)

		name, _, _ := strings.Cut(f.Tag.Get("json"), ",")
		if name == "-" {
			continue
		}

		if f.Anonymous && name == "" {
			if ft := deref(f.Type); ft.Kind() == reflect.Struct {
				out = append(out, describeArgs(ft)...)
				continue
			}
		}

		if !f.IsExported() {
			continue
		}

		if name == "" {
			name = f.Name
		}

		out = append(out, ArgMeta{
			Name:     name,
			Type:     argType(f.Type),
			Desc:     f.Tag.Get("desc"),
			Required: f.Tag.Get("required") == "true",
		})
	}

	return out
}

func argType(t reflect.Type) ArgType {
	t = deref(t)

	switch t.Kind() {
	case reflect.String:
		return ArgString
	case reflect.Bool:
		return ArgBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return ArgInt
	case reflect.Float32, reflect.Float64:
		return ArgFloat
	case reflect.Slice, reflect.Array:
		if t.Elem().Kind() == reflect.Uint8 {
			return ArgString
		}
		return ArgArray
	case reflect.Map, reflect.Struct:
		return ArgObject
	default:
		return ArgAny
	}
}

func deref(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

// decodeArgs decodes raw into A. Missing args are treated as an empty object,
// and unknown fields are rejected so typos surface as errors.
func decodeArgs[A any](raw json.RawMessage, required []string) (A, error) {
	var args A

	if len(raw) == 0 || string(raw) == "null" {
		raw = json.RawMessage("{}")
	}

	var present map[string]json.RawMessage
	if err := json.Unmarshal(raw, &present); err != nil {
		return args, fmt.Errorf("args must be a JSON object: %w", err)
	}

	for _, name := range required {
		if _, ok := present[name]; !ok {
			return args, fmt.Errorf("missing required argument %q", name)
		}
	}

	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&args); err != nil {
		return args, fmt.Errorf("invalid args: %w", err)
	}

	return args, nil
}
