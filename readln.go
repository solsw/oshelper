package oshelper

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// singleByteReader keeps a temporary buffer from consuming the next line.
type singleByteReader struct{ io.Reader }

func (r singleByteReader) Read(p []byte) (int, error) {
	return r.Reader.Read(p[:min(len(p), 1)])
}

// Freadln reads from 'r' until the first occurrence of end-of-line marker (see
// [bufio.ScanLines]), returning a string stripped of any trailing end-of-line marker.
// Returns [io.EOF] if 'r' is at EOF before any data is read.
// For efficient repeated reads, pass the same [bufio.Reader] on each call.
func Freadln(r io.Reader) (string, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(singleByteReader{r})
	}
	line, err := br.ReadString('\n')
	if err == io.EOF && len(line) > 0 {
		err = nil
	}
	line = strings.TrimSuffix(line, "\n")
	return strings.TrimSuffix(line, "\r"), err
}

// Readln is like [Freadln], but reads from standard input.
func Readln() (string, error) {
	return Freadln(os.Stdin)
}
