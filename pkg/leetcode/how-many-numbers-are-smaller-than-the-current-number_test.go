package leetcode

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{8, 1, 2, 2, 3},
			},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			name: "2",
			args: args{
				nums: []int{6, 5, 4, 8},
			},
			want: []int{2, 1, 0, 3},
		},
		{
			name: "3",
			args: args{
				nums: []int{7, 7, 7, 7},
			},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "3",
			args: args{
				nums: []int{37, 64, 63, 2, 41, 78, 51, 36, 2, 20, 25, 41, 72, 100, 17, 43, 54, 27, 34, 86, 12, 48, 70, 44, 87, 68, 62, 98, 68, 30, 8, 92, 5, 10},
			},
			want: []int{13, 23, 22, 0, 14, 28, 19, 12, 0, 7, 8, 14, 27, 33, 6, 16, 20, 9, 11, 29, 5, 18, 26, 17, 30, 24, 21, 32, 24, 10, 3, 31, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %#v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent2() = %v, want %v", got, tt.want)
			}
		})
	}
}
