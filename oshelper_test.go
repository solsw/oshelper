package oshelper

import (
	"testing"
)

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
