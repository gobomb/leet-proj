package panicRecover

import "testing"

func TestWt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"testpanic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Wt()
		})
	}
}
