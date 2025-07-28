package oshelper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/solsw/mathrandhelper"
)

// FileExistsFunc reports whether a regular file 'filename' exists.
// 'f' (if not nil) is used to process 'filename' before own error returning
// (e.g. 'f' may extract just file name from the full path).
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

// DirExistsFunc reports whether a directory 'dirname' exists.
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

// ClearDir clears the contents of the directory 'dirname'.
func ClearDir(dirname string) error {
	des, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, de := range des {
		if err := os.RemoveAll(filepath.Join(dirname, de.Name())); err != nil {
			return err
		}
	}
	return nil
}

// RandomOverFile overwrites the contents of file 'filename' with random data.
func RandomOverFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	const N int64 = 1024 * 1024
	quotient := fi.Size() / N
	remainder := fi.Size() % N
	bb := make([]byte, N)
	for range quotient {
		mathrandhelper.RandomBytes(bb)
		if _, err := f.Write(bb); err != nil {
			return err
		}
	}
	bb = make([]byte, remainder)
	mathrandhelper.RandomBytes(bb)
	if _, err := f.Write(bb); err != nil {
		return err
	}
	return f.Sync()
}

// WipeFile first overwrites the contents of file 'filename' with random data,
// then removes the file.
func WipeFile(filename string) error {
	if err := RandomOverFile(filename); err != nil {
		return err
	}
	return os.Remove(filename)
}
