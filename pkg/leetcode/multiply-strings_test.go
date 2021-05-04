package leetcode

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				"2", "3",
			},
			"6",
		},
		{
			"2",
			args{
				"123", "456",
			},
			"56088",
		},
		{
			"3",
			args{
				"123", "9",
			},
			"1107",
		},
		{
			"4",
			args{
				"123", "11",
			},
			"1353",
		},
		{
			"5",
			args{
				"123", "1198",
			},
			"147354",
		},
		{
			"6",
			args{
				"11", "100",
			},
			"1100",
		},
		{
			"7",
			args{
				"9133", "0",
			},
			"0",
		},
		{
			"7",
			args{
				"0", "0",
			},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
