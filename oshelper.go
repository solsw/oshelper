package oshelper

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// TempFileBase returns just a name of a temporary file.
// (See [os.CreateTemp] for 'pattern' usage.)
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

// ErrStdinNotRedirected is returned by StdinRedirected when [stdin] is not redirected.
//
// [stdin]: https://pkg.go.dev/os#pkg-variables
var ErrStdinNotRedirected = errors.New("standard input is not redirected (file '<' or pipe '|')")

// StdinRedirected returns nil if [stdin] is redirected. Otherwise, returns [ErrStdinNotRedirected].
//
// [stdin]: https://pkg.go.dev/os#pkg-variables
func StdinRedirected() error {
	finfo, _ := os.Stdin.Stat()
	fmode := finfo.Mode()
	if !(fmode.IsRegular() || (fmode&fs.ModeNamedPipe != 0)) {
		return ErrStdinNotRedirected
	}
	return nil
}
