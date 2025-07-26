package oshelper

import (
	"bufio"
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
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ss = append(ss, sc.Text())
	}
	return ss, sc.Err()
}

// WriteFileStringsNewLine writes 'ss' to the named file.
// Each string (including the last one) is followed by 'newLine'.
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
