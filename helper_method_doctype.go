package ace

import (
	"fmt"
	"io"
)

// Doctypes
var doctypes = map[string]string{
	"html":         `<!DOCTYPE html>`,
	"xml":          `<?xml version="1.0" encoding="utf-8" ?>`,
	"transitional": `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`,
	"strict":       `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`,
	"frameset":     `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`,
	"1.1":          `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
	"basic":        `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN" "http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">`,
	"mobile":       `<!DOCTYPE html PUBLIC "-//WAPFORUM//DTD XHTML Mobile 1.2//EN" "http://www.openmobilealliance.org/tech/DTD/xhtml-mobile12.dtd">`,
}

// helperMethodDoctype represents a helper method doctype.
type helperMethodDoctype struct {
	elementBase
	doctype string
}

// WriteTo writes data to w.
func (e *helperMethodDoctype) WriteTo(w io.Writer) (int64, error) {
	i, err := w.Write([]byte(doctypes[e.doctype]))
	return int64(i), err
}

// newHelperMethodDoctype creates and returns a helper method doctype.
func newHelperMethodDoctype(ln *line, src *source, parent element, opts *Options) (*helperMethodDoctype, error) {
	if len(ln.tokens) < 3 {
		return nil, fmt.Errorf("doctype is not specified [line: %d]", ln.no)
	}

	doctype := ln.tokens[2]

	if _, ok := doctypes[doctype]; !ok {
		return nil, fmt.Errorf("doctype is invalid [line: %d][doctype: %s]", ln.no, doctype)
	}

	e := &helperMethodDoctype{
		elementBase: newElementBase(ln, src, parent, opts),
		doctype:     doctype,
	}

	return e, nil
}
