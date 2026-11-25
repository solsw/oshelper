package oshelper

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestFreadln(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "empty",
			args:    args{r: bufio.NewReader(strings.NewReader(""))},
			want:    "",
			wantErr: true,
		},
		{name: "empty\r\n",
			args: args{r: bufio.NewReader(strings.NewReader("\r\n"))},
			want: "",
		},
		{name: "qwerty",
			args: args{r: bufio.NewReader(strings.NewReader("qwerty"))},
			want: "qwerty",
		},
		{name: "qwerty\n",
			args: args{r: bufio.NewReader(strings.NewReader("qwerty\n"))},
			want: "qwerty",
		},
		{name: "qwerty\r\nasdfgh\r\n",
			args: args{r: bufio.NewReader(strings.NewReader("qwerty\r\nasdfgh\r\n"))},
			want: "qwerty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Freadln(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Freadln() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Freadln() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreadlnSuccessiveLines(t *testing.T) {
	const input = "first\r\n\nlast\r"
	for _, r := range []io.Reader{strings.NewReader(input), bufio.NewReader(strings.NewReader(input))} {
		for _, want := range []string{"first", "", "last"} {
			got, err := Freadln(r)
			if got != want || err != nil {
				t.Errorf("Freadln(%T) = %q, %v; want %q, nil", r, got, err, want)
				break
			}
		}
		if got, err := Freadln(r); got != "" || err != io.EOF {
			t.Errorf("Freadln(%T) at EOF = %q, %v", r, got, err)
		}
	}
}

func TestFreadlnLongLine(t *testing.T) {
	want := strings.Repeat("x", 100_000)
	got, err := Freadln(strings.NewReader(want + "\n"))
	if err != nil || got != want {
		t.Errorf("Freadln returned %d bytes, %v; want %d bytes", len(got), err, len(want))
	}
}
