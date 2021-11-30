package medium

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
	}

	toTest := []func([]int) int{rob, rob1}

	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); got != tt.want {
					t.Errorf("%v(%+v) = %+v, want %+v", tt.args, funcName(f), got, tt.want)
				}
			})
		}
	}
}
