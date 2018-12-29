package gopass

import (
	"fmt"
	"io"
	"os"
)

type CleartextReader struct {
	Prompt string
	R      FdReader
	W      io.Writer
}

func NewCleartextReader() *CleartextReader {
	return &CleartextReader{
		Prompt: ">>> ",
		R:      os.Stdin,
		W:      os.Stdout,
	}
}

// SetPrompt changes the user prompt to s.
func (c *CleartextReader) SetPrompt(s string) {
	c.Prompt = s
}

// ReadSlice returns a line of input read from terminal, interruptable by ctrl-c.
// If prompt is not empty, it will be output as a prompt to the user
func (c *CleartextReader) ReadSlice() ([]byte, error) {
	var err error
	var line []byte
	bs := []byte("\b \b")
	r := c.R
	w := c.W

	if isTerminal(r.Fd()) {
		if oldState, err := makeRaw(r.Fd()); err != nil {
			return line, err
		} else {
			defer func() {
				restore(r.Fd(), oldState)
				fmt.Fprintln(w)
			}()
		}
	}

	if c.Prompt != "" {
		fmt.Fprint(w, c.Prompt)
	}

	// Track total bytes read, not just bytes in the password.  This ensures any
	// errors that might flood the console with nil or -1 bytes infinitely are
	// capped.
	var counter int
	for counter = 0; counter <= maxLength; counter++ {
		if v, e := getch(r); e != nil {
			err = e
			break
		} else if v == 127 || v == 8 {
			if l := len(line); l > 0 {
				line = line[:l-1]
				fmt.Fprint(w, string(bs))
			}
		} else if v == 13 || v == 10 {
			break
		} else if v == 3 {
			err = ErrInterrupted
			break
		} else if v != 0 {
			line = append(line, v)
			fmt.Fprint(w, string(v))
		}
	}

	if counter > maxLength {
		err = ErrMaxLengthExceeded
	}

	return line, err
}
