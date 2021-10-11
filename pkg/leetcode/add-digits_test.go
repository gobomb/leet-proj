package leetcode

import "testing"

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				num: 38,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				num: 0,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				num: 998,
			},
			want: 8,
		},
		{
			name: "4",
			args: args{
				num: 100,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
