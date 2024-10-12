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
			args: args{r: bufio.NewReader(strings.NewReader(""))},
			want: "",
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
