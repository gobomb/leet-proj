package easy

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		rs   []int
	}{
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			rs: []int{1, 3, 12, 0, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{0},
			},
			rs: []int{0},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0, 3, 12, 0, 0, 0, 0, 1},
			},
			rs: []int{1, 3, 12, 1, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.rs) {
				t.Fatalf("failed want %v got %v", tt.rs, tt.args.nums)
			}
		})
	}
}
