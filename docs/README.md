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

Opens `dirname` once, then lists and recursively removes its entries through
the same directory handle. The directory itself stays in place.

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

`Freadln` reads one line per call and strips the LF or CRLF ending. It returns
a final unterminated line with a nil error, stripping a trailing carriage
return if present. A blank line returns `("", nil)`. A call at end of input
returns `("", io.EOF)`.

Successive calls on the same reader preserve the remaining lines, including
lines longer than 64 KiB. For efficient reads in a loop, create a
`*bufio.Reader` once and pass it to every call. `Readln` reads one line at a
time from standard input with the same line-ending and EOF behavior.

### ReadFileStrings

```go
func ReadFileStrings(filename string) ([]string, error)
```

Returns one string per line using the same line-ending rules as `Freadln`.
It handles lines longer than 64 KiB and includes a final unterminated line.
An empty file returns a nil slice and a nil error.

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

`SyncFileWriter` implements `io.Writer` by writing to the underlying file and
calling `Sync` after each successful write. If the write succeeds and `Sync`
fails, `Write` returns the byte count from the write together with the sync
error. In that case the file accepted those bytes, but their durability is
uncertain. `File` returns the underlying `*os.File`.

## Standard input

### StdinRedirected, ErrStdinNotRedirected

```go
var ErrStdinNotRedirected = errors.New("standard input is not redirected (file '<' or pipe '|')")

func StdinRedirected() error
```

`StdinRedirected` returns nil if standard input is redirected (from a file via
`<` or a pipe via `|`). Otherwise, it returns `ErrStdinNotRedirected`.
