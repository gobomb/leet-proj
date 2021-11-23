package medium

import (
	"testing"
)

func Test_findMin(t *testing.T) {
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
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				nums: []int{11, 13, 15, 17},
			},
			want: 11,
		},
		{
			name: "4",
			args: args{
				nums: []int{13, 11},
			},
			want: 11,
		},
		{
			name: "5",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "6",
			args: args{
				nums: []int{31, 41, 1, 2, 3},
			},
			want: 1,
		},
	}

	totests := []func([]int) int{findMin, findMin1}

	for _, f := range totests {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); got != tt.want {
					t.Errorf("%v() = %#v, want %#v", funcName(f), got, tt.want)
				}
			})
		}
	}
}

func Test_findMinII(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "升序无重复",
			args: args{
				nums: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "旋转无重复",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "旋转重复",
			args: args{
				nums: []int{12, 12, 12, 12, 12, 12, 13, 15, 17, 11},
			},
			want: 11,
		},
		{
			name: "4",
			args: args{
				nums: []int{13, 12, 12},
			},
			want: 12,
		},
		{
			name: "第一和倒数第一相等，最小值在右半部分",
			args: args{
				nums: []int{3, 3, 3, 1, 3},
			},
			want: 1,
		},
		{
			name: "第一和倒数第一相等，最小值在左半部分",
			args: args{
				nums: []int{3, 1, 3, 3, 3},
			},
			want: 1,
		},
		{
			name: "6",
			args: args{
				nums: []int{1, 1, 3, 1},
			},
			want: 1,
		},
		{
			name: "7",
			args: args{
				nums: []int{1, 3, 3},
			},
			want: 1,
		},
	}

	totests := []func([]int) int{findMinII, findMinII1}

	for _, f := range totests {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); got != tt.want {
					t.Errorf("%v() = %#v, want %#v", funcName(f), got, tt.want)
				}
			})
		}
	}
}
