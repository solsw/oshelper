package oshelper

import (
	"bufio"
	"io"
	"os"
)

// Freadln reads from 'r' until the first occurrence of end-of-line marker (see
// [bufio.ScanLines]), returning a string stripped of any trailing end-of-line marker.
func Freadln(r io.Reader) (string, error) {
	bs := bufio.NewScanner(r)
	ok := bs.Scan()
	if !ok {
		return "", bs.Err()
	}
	return bs.Text(), nil
}

// Readln is like [Freadln], but reads from standard input.
func Readln() (string, error) {
	return Freadln(os.Stdin)
}
