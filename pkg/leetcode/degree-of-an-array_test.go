package leetcode

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 2, 3, 1, 4, 2},
			},
			want: 6,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 1, 2, 2, 2, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
