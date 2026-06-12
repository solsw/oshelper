# oshelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/oshelper.svg)](https://pkg.go.dev/github.com/solsw/oshelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/oshelper)

Package `oshelper` contains helpers for Go's [os](https://pkg.go.dev/os) package.

```text
go get github.com/solsw/oshelper
```

```go
import "github.com/solsw/oshelper"
```

## Constants

### NewLine

```go
const NewLine = "\r\n" // on Windows
const NewLine = "\n"   // on other platforms
```

Platform-specific end-of-line marker, selected at build time.

## Files and directories

### FileExists, FileExistsFunc

```go
func FileExists(filename string) (bool, error)
func FileExistsFunc(filename string, f func(string) string) (bool, error)
```

Report whether a regular file `filename` exists. An empty `filename` is an error.
If `filename` exists but is not a regular file, an error is returned; in
`FileExistsFunc`, `f` (if not nil) processes `filename` before it is embedded
in that error (e.g. `f` may extract just the file name from the full path).

### DirExists, DirExistsFunc

```go
func DirExists(dirname string) (bool, error)
func DirExistsFunc(dirname string, f func(string) string) (bool, error)
```

Report whether a directory `dirname` exists. An empty `dirname` is an error.
If `dirname` exists but is not a directory, an error is returned; in
`DirExistsFunc`, `f` (if not nil) processes `dirname` before it is embedded
in that error (e.g. `f` may shorten an excessively long `dirname`).

### ClearDir

```go
func ClearDir(dirname string) error
```

Removes the contents of the directory `dirname`, leaving the directory itself in place.

### RandomOverFile

```go
func RandomOverFile(filename string) error
```

Overwrites the contents of the file `filename` with cryptographically random
data of the same size, then syncs the file to stable storage.

### WipeFile

```go
func WipeFile(filename string) error
```

First overwrites the contents of the file `filename` with random data
(see `RandomOverFile`), then removes the file.

### TempFileBase

```go
func TempFileBase(pattern string) (string, error)
```

Returns just the name (without directory) of a temporary file. The file itself
is created and immediately removed. See [os.CreateTemp](https://pkg.go.dev/os#CreateTemp)
for `pattern` usage.

### ExeDir

```go
func ExeDir() (string, error)
```

Returns an absolute representation of the directory name of the executable
that started the current process.

## Environment variables

### GetenvDef

```go
func GetenvDef(key, def string) string
```

Retrieves the value of the environment variable named by `key`.
If the variable is not present or has an empty value, `def` is returned.

### GetenvErr

```go
func GetenvErr(key string) (string, error)
```

Retrieves the value of the environment variable named by `key`.
If the variable is not present or has an empty value, an error is returned.

## Reading and writing

### Freadln, Readln

```go
func Freadln(r io.Reader) (string, error)
func Readln() (string, error)
```

`Freadln` reads from `r` until the first end-of-line marker (see
[bufio.ScanLines](https://pkg.go.dev/bufio#ScanLines)) and returns the line
stripped of any trailing end-of-line marker. Returns `io.EOF` if `r` is at EOF
before any data is read. `Readln` is like `Freadln`, but reads from standard input.

### ReadFileStrings

```go
func ReadFileStrings(filename string) ([]string, error)
```

Returns the contents of the file `filename` as a `[]string`, one element per line.

### WriteFileStrings, WriteFileStringsNewLine

```go
func WriteFileStrings(filename string, ss []string, perm os.FileMode) error
func WriteFileStringsNewLine(filename string, ss []string, perm os.FileMode, newLine string) error
```

Write `ss` to the named file. Each string (including the last one) is followed
by `newLine` (`WriteFileStringsNewLine`) or by `NewLine` (`WriteFileStrings`).
If `ss` is empty, only the newline is written. See
[os.WriteFile](https://pkg.go.dev/os#WriteFile) for `perm` usage.

### SyncFileWriter

```go
type SyncFileWriter struct { /* unexported fields */ }

func NewSyncFileWriter(f *os.File) *SyncFileWriter
func (sfw *SyncFileWriter) File() *os.File
func (sfw *SyncFileWriter) Write(p []byte) (n int, err error)
```

An [os.File](https://pkg.go.dev/os#File)-based
[io.Writer](https://pkg.go.dev/io#Writer) implementation that calls
[os.File.Sync](https://pkg.go.dev/os#File.Sync) after each write, so every
write is committed to stable storage. `File` returns the underlying file.

## Standard input

### StdinRedirected, ErrStdinNotRedirected

```go
var ErrStdinNotRedirected = errors.New("standard input is not redirected (file '<' or pipe '|')")

func StdinRedirected() error
```

`StdinRedirected` returns nil if standard input is redirected (from a file via
`<` or a pipe via `|`). Otherwise, it returns `ErrStdinNotRedirected`.
