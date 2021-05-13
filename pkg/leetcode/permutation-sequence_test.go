package leetcode

import "testing"

func Test_getPermutation(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				3, 3,
			},
			"213",
		},
		{
			"2",
			args{
				4, 9,
			},
			"2314",
		},
		{
			"3",
			args{
				3, 1,
			},
			"123",
		},
		{
			"4",
			args{
				1, 1,
			},
			"1",
		},
		{
			"5",
			args{
				0, 0,
			},
			"",
		},
		{
			"6",
			args{
				2, 2,
			},
			"21",
		},
		{
			"8",
			args{
				3, 2,
			},
			"132",
		},
		// {
		// 	"7",
		// 	args{
		// 		9, 233794,
		// 	},
		// 	"0",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
