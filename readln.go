package oshelper

import (
	"bufio"
	"io"
	"os"
)

// Freadln reads from 'r' until the first occurrence of end-of-line marker (see
// [bufio.ScanLines]), returning a string stripped of any trailing end-of-line marker.
// Returns [io.EOF] if 'r' is at EOF before any data is read.
func Freadln(r io.Reader) (string, error) {
	bs := bufio.NewScanner(r)
	ok := bs.Scan()
	if !ok {
		if err := bs.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return bs.Text(), nil
}

// Readln is like [Freadln], but reads from standard input.
func Readln() (string, error) {
	return Freadln(os.Stdin)
}
