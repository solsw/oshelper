package oshelper

import (
	"os"
	"path/filepath"
	"testing"
)

// added by Claude Code

func TestReadFileStrings(t *testing.T) {
	td := t.TempDir()
	emptyPath := filepath.Join(td, "empty.txt")
	_ = os.WriteFile(emptyPath, []byte{}, os.ModePerm)
	multiPath := filepath.Join(td, "multi.txt")
	_ = os.WriteFile(multiPath, []byte("line1\nline2\nline3"), os.ModePerm)
	tests := []struct {
		name     string
		filename string
		want     []string
		wantErr  bool
	}{
		{name: "nonexistent",
			filename: filepath.Join(td, "missing.txt"),
			wantErr:  true,
		},
		{name: "empty",
			filename: emptyPath,
			want:     nil,
		},
		{name: "multi",
			filename: multiPath,
			want:     []string{"line1", "line2", "line3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileStrings(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ReadFileStrings() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ReadFileStrings()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestWriteFileStringsNewLine(t *testing.T) {
	td := t.TempDir()
	tests := []struct {
		name     string
		ss       []string
		newLine  string
		wantFile string
	}{
		{name: "empty",
			ss:       nil,
			newLine:  "\n",
			wantFile: "\n",
		},
		{name: "single",
			ss:       []string{"hello"},
			newLine:  "\n",
			wantFile: "hello\n",
		},
		{name: "multi",
			ss:       []string{"a", "b", "c"},
			newLine:  "\n",
			wantFile: "a\nb\nc\n",
		},
		{name: "crlf",
			ss:       []string{"x", "y"},
			newLine:  "\r\n",
			wantFile: "x\r\ny\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(td, tt.name+".txt")
			if err := WriteFileStringsNewLine(path, tt.ss, os.ModePerm, tt.newLine); err != nil {
				t.Errorf("WriteFileStringsNewLine() error = %v", err)
				return
			}
			got, _ := os.ReadFile(path)
			if string(got) != tt.wantFile {
				t.Errorf("WriteFileStringsNewLine() file = %q, want %q", got, tt.wantFile)
			}
		})
	}
}

func TestWriteFileStrings(t *testing.T) {
	td := t.TempDir()
	path := filepath.Join(td, "out.txt")
	ss := []string{"line1", "line2"}
	if err := WriteFileStrings(path, ss, os.ModePerm); err != nil {
		t.Errorf("WriteFileStrings() error = %v", err)
		return
	}
	got, _ := os.ReadFile(path)
	want := "line1" + NewLine + "line2" + NewLine
	if string(got) != want {
		t.Errorf("WriteFileStrings() file = %q, want %q", got, want)
	}
}
