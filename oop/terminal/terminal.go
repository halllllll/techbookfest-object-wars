package terminal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Terminal struct {
	in_stream  *bufio.Reader
	out_stream io.Writer
}

func New(r io.Reader, w io.Writer) *Terminal {
	return &Terminal{
		in_stream:  bufio.NewReader(r),
		out_stream: w,
	}
}

func (t *Terminal) Prompt(prompt string) (string, error) {
	if _, err := fmt.Fprint(t.out_stream, prompt); err != nil {
		return "", err
	}
	line, err := t.in_stream.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(line, "\r\n"), nil
}

func (t *Terminal) Print(output string) error {
	_, err := fmt.Fprint(t.out_stream, output)
	return err
}

func (t *Terminal) EmptyLine() error {
	return t.Print("\n")
}
