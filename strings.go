package oshelper

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadFileStrings returns contents of the file 'filename' as []string.
func ReadFileStrings(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var ss []string
	br := bufio.NewReader(f)
	for {
		line, err := Freadln(br)
		if err == io.EOF {
			return ss, nil
		}
		if err != nil {
			if line != "" {
				ss = append(ss, line)
			}
			return ss, err
		}
		ss = append(ss, line)
	}
}

// WriteFileStringsNewLine writes 'ss' to the named file.
// Each string (including the last one) is followed by 'newLine'.
// If 'ss' is empty, only 'newLine' is written to the file.
// (See [os.WriteFile] for 'perm' usage.)
func WriteFileStringsNewLine(filename string, ss []string, perm os.FileMode, newLine string) error {
	return os.WriteFile(filename, []byte(strings.Join(ss, newLine)+newLine), perm)
}

// WriteFileStrings writes 'ss' to the named file.
// Each string (including the last one) is followed by [NewLine].
// (See [os.WriteFile] for 'perm' usage.)
func WriteFileStrings(filename string, ss []string, perm os.FileMode) error {
	return WriteFileStringsNewLine(filename, ss, perm, NewLine)
}
