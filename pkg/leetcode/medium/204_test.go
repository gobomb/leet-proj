package medium

import "testing"

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 10,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 0,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		n: 499979,
		// 	},
		// 	want: 41537,
		// },
		{
			name: "",
			args: args{
				n: 3,
			},
			want: 1,
		},
	}

	toTest := []func(int) int{countPrimes2}

	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.n); got != tt.want {
					t.Errorf("%v() = %+v, want %+v", funcName(f), got, tt.want)
				}
			})
		}
	}
}
