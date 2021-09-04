package leetcode

import "testing"

func Test_judgeCircle(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				moves: "UD",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				moves: "LL",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				moves: "RRDD",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				moves: "LDRRLRUULR",
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				moves: "RLUURDDDLU",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.args.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
