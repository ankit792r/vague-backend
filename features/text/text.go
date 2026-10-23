package text

import "bytes"

// Offset is a byte offset from the start of the buffer.
type Offset int

type Text struct {
	lines  [][]byte
	starts []Offset
}

// New builds a Text from raw bytes. b is copied and not retained.
func New(b []byte) *Text {
	t := &Text{}
	t.SetBytes(b)
	return t
}

// SetBytes replaces the entire contents.
func (t *Text) SetBytes(b []byte) {
	t.lines = bytes.Split(bytes.Clone(b), []byte{'\n'})
	t.starts = nil
}
