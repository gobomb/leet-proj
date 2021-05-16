package leetcode

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				"11",
				"1",
			},
			"100",
		},
		{
			"2",
			args{
				"1010",
				"1011",
			},
			"10101",
		},
		{
			"3",
			args{
				"1011",
				"1011",
			},
			"10110",
		},
		{
			"5",
			args{
				"1",
				"0",
			},
			"1",
		},
		{
			"6",
			args{
				"11111",
				"1",
			},
			"100000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
