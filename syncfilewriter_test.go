package oshelper

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSyncFileWriter(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sync.txt")
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	sfw := NewSyncFileWriter(f)
	if sfw.File() != f {
		t.Errorf("SyncFileWriter.File() = %v, want %v", sfw.File(), f)
	}
	data := []byte("hello")
	n, err := sfw.Write(data)
	if err != nil {
		t.Errorf("SyncFileWriter.Write() error = %v", err)
	}
	if n != len(data) {
		t.Errorf("SyncFileWriter.Write() n = %v, want %v", n, len(data))
	}
	f.Close()
	got, _ := os.ReadFile(path)
	if string(got) != "hello" {
		t.Errorf("SyncFileWriter content = %q, want %q", got, "hello")
	}
}

func TestSyncFileWriterCountOnSyncError(t *testing.T) {
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if f.Sync() == nil {
		t.Skip("null device supports Sync on this platform")
	}
	data := []byte("hello")
	n, err := NewSyncFileWriter(f).Write(data)
	if n != len(data) || err == nil {
		t.Errorf("Write = %d, %v; want %d and a sync error", n, err, len(data))
	}
}
