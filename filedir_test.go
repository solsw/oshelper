package oshelper

import (
	"os"
	"path/filepath"
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

func TestClearDir(t *testing.T) {
	const dir1 = "dir1"
	const dir2 = "dir2"
	_ = os.Mkdir(dir1, os.ModePerm)
	defer os.Remove(dir1)
	_ = os.WriteFile(filepath.Join(dir1, "file1"), []byte("qwerty"), os.ModePerm)
	_ = os.Mkdir(filepath.Join(dir1, dir2), os.ModePerm)
	_ = os.WriteFile(filepath.Join(dir1, dir2, "file2"), []byte("qwerty"), os.ModePerm)

	type args struct {
		dirname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1", args: args{dirname: dir1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ClearDir(tt.args.dirname); (err != nil) != tt.wantErr {
				t.Errorf("ClearDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
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
