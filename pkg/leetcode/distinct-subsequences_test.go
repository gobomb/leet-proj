package leetcode

import "testing"

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "rabbbit",
				t: "rabbit",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				s: "babgbag",
				t: "bag",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct2() = %v, want %v", got, tt.want)
			}
		})
	}
}
