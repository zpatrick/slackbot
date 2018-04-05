package slackbot

import (
	"fmt"
	"io"
)

// Write writes input to w
func Write(w io.Writer, input []byte) error {
	if _, err := w.Write(input); err != nil {
		return err
	}

	return nil
}

// WriteString writes text to w
func WriteString(w io.Writer, text string) error {
	return Write(w, []byte(text))
}

// WriteStringf writes the formatted text to w.
func WriteStringf(w io.Writer, format string, tokens ...interface{}) error {
	return WriteString(w, fmt.Sprintf(format, tokens...))
}

// The WriterFunc type is an adapter to allow the use of ordinary functions as io.Writers.
type WriterFunc func(p []byte) (n int, err error)

// Write calls f(p)
func (f WriterFunc) Write(p []byte) (n int, err error) {
	return f(p)
}
