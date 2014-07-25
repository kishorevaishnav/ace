package ace

import (
	"bytes"
	"io"
	"strings"
)

// action represents an action.
type action struct {
	elementBase
}

// Do nothing.
func (e *action) WriteTo(w io.Writer) (int64, error) {
	var bf bytes.Buffer

	// Write the action
	bf.WriteString(strings.TrimSpace(e.ln.str))

	// Write the children's HTML.
	if i, err := e.writeChildren(&bf); err != nil {
		return i, err
	}

	// Write the buffer.
	i, err := w.Write(bf.Bytes())

	return int64(i), err

}

// newAction creates and returns an action.
func newAction(ln *line, src *source, parent element, opts *Options) *action {
	return &action{
		elementBase: newElementBase(ln, src, parent, opts),
	}
}
