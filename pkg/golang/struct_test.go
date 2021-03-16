package golang

import "testing"

func Test_kk_test(t *testing.T) {
	tests := []struct {
		name string
		k    kk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.test()
		})
	}
}

func Test_size(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size()
		})
	}
}
