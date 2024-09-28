package oshelper

import (
	"testing"
)

func TestGetenvDef(t *testing.T) {
	type args struct {
		key string
		def string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				key: "C3043E18D2234F2897BE0BCEBBE0C840",
				def: "qwerty",
			},
			want: "qwerty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetenvDef(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetenvDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetenvErr(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "PATH",
			args: args{
				key: "PATH",
			},
			wantErr: false,
		},
		{name: "C3043E18D2234F2897BE0BCEBBE0C840",
			args: args{
				key: "C3043E18D2234F2897BE0BCEBBE0C840",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetenvErr(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetenvErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
