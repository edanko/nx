package jinfeng

import (
	"testing"
)

func TestSpec(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"PP100*6.0", "P0100x06.0", false},
		{"PP120*6.5", "P0120x06.5", false},
		{"PP140*7.0", "P0140x07.0", false},
		{"PP140*9.0", "P0140x09.0", false},
		{"PP160*8.0", "P0160x08.0", false},
		{"PP160*10.0", "P0160x10.0", false},
		{"PP180*9.0", "P0180x09.0", false},
		{"PP180*11.0", "P0180x11.0", false},
		{"PP200*10.0", "P0200x10.0", false},
		{"PP200*11.0", "P0200x11.0", false},
		{"PP200*12.0", "P0200x12.0", false},
		{"PP220*11.0", "P0220x11.0", false},
		{"PP220*13.0", "P0220x13.0", false},
		{"PP240*12.0", "P0240x12.0", false},
		{"PP240*14.0", "P0240x14.0", false},
		{"PP999*99.0", "", true},
		{"HZ999*99.0", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Spec(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Spec() = %v, want %v", got, tt.want)
			}
		})
	}
}
