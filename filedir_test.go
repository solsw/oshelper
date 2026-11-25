package oshelper

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestFileExists(t *testing.T) {
	f, _ := os.CreateTemp("", "")
	f.Close()
	defer os.Remove(f.Name())
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "1e",
			args: args{
				filename: "",
			},
			want:    false,
			wantErr: true,
		},
		{name: "1",
			args: args{
				filename: "C3043E18D2234F2897BE0BCEBBE0C840",
			},
			want: false,
		},
		{name: "2",
			args: args{
				filename: f.Name(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileExists(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirExists(t *testing.T) {
	td := t.TempDir()
	f, _ := os.CreateTemp("", "")
	f.Close()
	defer os.Remove(f.Name())
	tests := []struct {
		name    string
		dirname string
		want    bool
		wantErr bool
	}{
		{name: "empty", dirname: "", want: false, wantErr: true},
		{name: "nonexistent", dirname: "C3043E18D2234F2897BE0BCEBBE0C840", want: false},
		{name: "existing", dirname: td, want: true},
		{name: "file", dirname: f.Name(), want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirExists(tt.dirname)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWipeFile(t *testing.T) {
	td := t.TempDir()
	filename := filepath.Join(td, "wipe.txt")
	_ = os.WriteFile(filename, []byte("secret data"), os.ModePerm)
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{name: "nonexistent", filename: filepath.Join(td, "missing.txt"), wantErr: true},
		{name: "existing", filename: filename},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WipeFile(tt.filename); (err != nil) != tt.wantErr {
				t.Errorf("WipeFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if _, err := os.Stat(tt.filename); !os.IsNotExist(err) {
					t.Errorf("WipeFile() file still exists: %s", tt.filename)
				}
			}
		})
	}
}

func TestWipeFileRejectsSymlink(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "target")
	link := filepath.Join(dir, "link")
	const contents = "keep this file"
	if err := os.WriteFile(target, []byte(contents), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("cannot create symlink: %v", err)
	}

	if err := WipeFile(link); err == nil {
		t.Error("WipeFile accepted a symlink")
	}
	if _, err := os.Readlink(link); err != nil {
		t.Errorf("symlink was not preserved: %v", err)
	}
	got, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != contents {
		t.Error("WipeFile modified the symlink target")
	}
}

func TestClearDir(t *testing.T) {
	dir := t.TempDir()
	subdir := filepath.Join(dir, "subdir")
	if err := os.Mkdir(subdir, 0700); err != nil {
		t.Fatal(err)
	}
	for _, parent := range []string{dir, subdir} {
		if err := os.WriteFile(filepath.Join(parent, "file"), nil, 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := ClearDir(dir); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("entries left after ClearDir: %v", entries)
	}
}

func TestClearDirThroughSymlinkParent(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows resolves .. before following directory symlinks")
	}
	base := t.TempDir()
	child := filepath.Join(base, "real", "child")
	work := filepath.Join(base, "real", "work")
	for _, dir := range []string{child, work} {
		if err := os.MkdirAll(dir, 0700); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(work, "file"), nil, 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(base, "link")
	if err := os.Symlink(child, link); err != nil {
		t.Skipf("cannot create symlink: %v", err)
	}
	// Keep .. in the input so the filesystem resolves it relative to child.
	if err := ClearDir(link + "/../work"); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(work)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("entries left in resolved directory: %v", entries)
	}
}

func TestRandomOverFile(t *testing.T) {
	testFilename, _ := TempFileBase("")
	testFilename = filepath.Join(t.TempDir(), testFilename)
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1",
			args:    args{filename: testFilename},
			wantErr: true,
		},
		{name: "2",
			args: args{filename: testFilename},
		},
		// {name: "3",
		// 	args: args{filename: "t.t"},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "2" {
				_ = os.WriteFile(testFilename, []byte("12345678\nqwertyui\nasdfghjk\nzxcvbnm,"), os.ModePerm)
			}
			if err := RandomOverFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("RandomOverFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
