package slackbot

import (
	"fmt"
	"io"
)

func Write(w io.Writer, input []byte) error {
	if _, err := w.Write(input); err != nil {
		return err
	}

	return nil
}

func WriteString(w io.Writer, text string) error {
	return Write(w, []byte(text))
}

func WriteStringf(w io.Writer, format string, tokens ...interface{}) error {
	return WriteString(w, fmt.Sprintf(format, tokens...))
}
