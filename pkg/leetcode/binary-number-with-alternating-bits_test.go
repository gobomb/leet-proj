package leetcode

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				n: 5,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				n: 7,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				n: 11,
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				n: 10,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				n: 3,
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				n: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBits(tt.args.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
