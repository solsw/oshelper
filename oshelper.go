package oshelper

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileExistsFunc reports whether a regular file 'filename' exists.
//
// 'f' (if not nil) is used to process 'filename' before own error returning
// (e.g. 'f' may extract just file name from full path).
func FileExistsFunc(filename string, f func(string) string) (bool, error) {
	if filename == "" {
		return false, errors.New("empty filename")
	}
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	// 'filename' exists
	if !fi.Mode().IsRegular() {
		if f != nil {
			filename = f(filename)
		}
		return false, fmt.Errorf("not a regular file '%s'", filename)
	}
	return true, nil
}

// FileExists reports whether a regular file 'filename' exists.
func FileExists(filename string) (bool, error) {
	return FileExistsFunc(filename, nil)
}

// FileExistsMust is like FileExists but returns 'false' in case of error.
func FileExistsMust(filename string) bool {
	fe, err := FileExists(filename)
	if err != nil {
		return false
	}
	return fe
}

// DirExistsFunc reports whether a directory 'dirname' exists.
//
// 'f' (if not nil) is used to process 'dirname' before own error returning
// (e.g. 'f' may shorten excessively long 'dirname').
func DirExistsFunc(dirname string, f func(string) string) (bool, error) {
	if dirname == "" {
		return false, errors.New("empty dirname")
	}
	fi, err := os.Stat(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	// 'dirname' exists
	if !fi.IsDir() {
		if f != nil {
			dirname = f(dirname)
		}
		return false, fmt.Errorf("not a directory '%s'", dirname)
	}
	return true, nil
}

// DirExists reports whether a directory 'dirname' exists.
func DirExists(dirname string) (bool, error) {
	return DirExistsFunc(dirname, nil)
}

// DirExistsMust is like DirExists but returns 'false' in case of error.
func DirExistsMust(dirname string) bool {
	de, err := DirExists(dirname)
	if err != nil {
		return false
	}
	return de
}

// TempFileBase returns just a name of a temporary file.
// (See os.CreateTemp for 'pattern' usage.)
func TempFileBase(pattern string) (string, error) {
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
	return filepath.Base(f.Name()), nil
}

// ExeDir returns an absolute representation of the directory name
// of the executable that started the current process.
func ExeDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir, err := filepath.Abs(filepath.Dir(exe))
	if err != nil {
		return "", err
	}
	return exeDir, nil
}

// GetenvDef retrieves the value of the environment variable named by the 'key'.
// If the value is empty or the variable is not present 'def' is returned.
func GetenvDef(key, def string) string {
	r := os.Getenv(key)
	if r == "" {
		return def
	}
	return r
}

// ReadFileStrings returns the contents of the file 'filename' as []string.
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
// (See os.WriteFile for details.)
func WriteFileStringsNewLine(filename string, ss []string, perm os.FileMode, newLine string) error {
	return os.WriteFile(filename, []byte(strings.Join(ss, newLine)+newLine), perm)
}

// WriteFileStrings writes 'ss' to the named file.
// Each string (including the last one) is followed by oshelper.NewLine.
// (See os.WriteFile for details.)
func WriteFileStrings(filename string, ss []string, perm os.FileMode) error {
	return WriteFileStringsNewLine(filename, ss, perm, NewLine)
}
