package oshelper

import (
	"path/filepath"
	"testing"
)

func TestTempFileBase(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		wantErr bool
	}{
		{name: "empty pattern", pattern: ""},
		{name: "with pattern", pattern: "test-*"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TempFileBase(tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("TempFileBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != filepath.Base(got) {
				t.Errorf("TempFileBase() returned path with directory: %v", got)
			}
		})
	}
}

func TestExeDir(t *testing.T) {
	got, err := ExeDir()
	if err != nil {
		t.Errorf("ExeDir() error = %v", err)
		return
	}
	if !filepath.IsAbs(got) {
		t.Errorf("ExeDir() returned non-absolute path: %v", got)
	}
}

func TestStdinRedirected(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "1", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StdinRedirected(); (err != nil) != tt.wantErr {
				t.Errorf("StdinRedirected() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
