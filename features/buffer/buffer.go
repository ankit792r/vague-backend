package buffer

import (
	"time"
	"vague/features/text"
)

type Buffer struct {
	ID   uint64
	Name string
	Path string

	Binary   bool
	ReadOnly bool

	Text *text.Text

	startedEmpty bool
	onDisk       bool
	diskSize     int64
	diskModTime  time.Time
	savedSeq     int
}

func NewScratch(id uint64, name string) *Buffer {
	initial := []byte("This is scratch buffer\nModified contents are not saved.")
	buf := &Buffer{
		ID:   id,
		Name: name,
		Text: text.New(initial),
		// History: text.NewUndoTree(),
	}
	// buf.History.SetInitial(initial)
	return buf
}
