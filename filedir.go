package oshelper

import (
	cryptorand "crypto/rand"
	"errors"
	"fmt"
	"io/fs"
	"os"
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
		if errors.Is(err, fs.ErrNotExist) {
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
		if errors.Is(err, fs.ErrNotExist) {
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
	root, err := os.OpenRoot(dirname)
	if err != nil {
		return err
	}
	defer root.Close()
	des, err := fs.ReadDir(root.FS(), ".")
	if err != nil {
		return err
	}
	for _, de := range des {
		if err := root.RemoveAll(de.Name()); err != nil {
			return err
		}
	}
	return nil
}

// RandomOverFile overwrites the contents of file 'filename' with random data.
func RandomOverFile(filename string) error {
	return randomOverFile(filename, nil)
}

func randomOverFile(filename string, expected fs.FileInfo) error {
	f, err := os.OpenFile(filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if expected != nil && !os.SameFile(expected, fi) {
		return fmt.Errorf("file changed before overwrite: %q", filename)
	}
	const N int64 = 1024 * 1024
	quotient := fi.Size() / N
	if quotient > 0 {
		bb := make([]byte, N)
		for range quotient {
			_, _ = cryptorand.Read(bb)
			if _, err := f.Write(bb); err != nil {
				return err
			}
		}
	}
	remainder := fi.Size() % N
	if remainder > 0 {
		bb := make([]byte, remainder)
		_, _ = cryptorand.Read(bb)
		if _, err := f.Write(bb); err != nil {
			return err
		}
	}
	return f.Sync()
}

// WipeFile overwrites a regular file with random data, then removes it.
// It rejects a symbolic link at the final path component.
// The caller must keep the path stable until WipeFile returns.
func WipeFile(filename string) error {
	fi, err := os.Lstat(filename)
	if err != nil {
		return err
	}
	if !fi.Mode().IsRegular() {
		return fmt.Errorf("not a regular file: %q", filename)
	}
	if err := randomOverFile(filename, fi); err != nil {
		return err
	}
	current, err := os.Lstat(filename)
	if err != nil {
		return err
	}
	if !current.Mode().IsRegular() || !os.SameFile(fi, current) {
		return fmt.Errorf("file changed before removal: %q", filename)
	}
	return os.Remove(filename)
}
