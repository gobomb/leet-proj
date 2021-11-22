package medium

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				left:  5,
				right: 7,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				left:  0,
				right: 0,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				left:  1,
				right: 2147483647,
			},
			want: 0,
		},
	}

	toTest := []func(int, int) int{rangeBitwiseAnd, rangeBitwiseAnd1}

	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.left, tt.args.right); got != tt.want {
					t.Errorf("%v() = %+v, want %+v", funcName(f), got, tt.want)
				}
			})
		}
	}
}
