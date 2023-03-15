package oshelper

import (
	"os"
)

// SyncFileWriter is the [os.File]-based [io.Writer] implementation,
// that calls [os.File.Sync] on the underlying file after each [os.File.Write] call.
type SyncFileWriter struct {
	file *os.File
}

// NewSyncFileWriter creates a new [SyncFileWriter] based on the file 'f'.
func NewSyncFileWriter(f *os.File) *SyncFileWriter {
	return &SyncFileWriter{file: f}
}

// Write implements the [io.Writer] interface.
func (sfw *SyncFileWriter) Write(p []byte) (n int, err error) {
	n, err = sfw.file.Write(p)
	if err != nil {
		return
	}
	err = sfw.file.Sync()
	return
}
