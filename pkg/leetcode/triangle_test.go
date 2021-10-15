package leetcode

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				triangle: [][]int{
					{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
				},
			},
			want: 11,
		},
		{
			name: "2",
			args: args{
				triangle: [][]int{
					{-10},
				},
			},
			want: -10,
		},
		{
			name: "3",
			args: args{
				triangle: [][]int{
					{-1}, {2, 3}, {1, -1, -3},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
