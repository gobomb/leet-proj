package panicRecover

import "testing"

func TestPr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Add test cases"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pr()
		})
	}
}

func Test_fakeFun(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeFun(tt.args.s)
		})
	}
}
