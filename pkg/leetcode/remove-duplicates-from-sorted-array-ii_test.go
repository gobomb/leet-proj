package leetcode

import (
	"justest/pkg/utils"
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_removeDuplicatesII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantInt []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want:    5,
			wantInt: []int{1, 1, 2, 2, 3},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want:    7,
			wantInt: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 0},
			},
			want:    2,
			wantInt: []int{0, 0},
		},
		{
			name: "4",
			args: args{
				nums: []int{0, 0, 0, 0, 0},
			},
			want:    2,
			wantInt: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := utils.DeepCopyIntSlice(tt.args.nums)
			if _ = removeDuplicatesII(tt.args.nums); !reflect.DeepEqual(tt.args.nums[:tt.want], tt.wantInt) {
				t.Errorf("removeDuplicatesII(%#v) = %#v, want %#v", before, tt.args.nums[:tt.want], tt.wantInt)
			}
		})
	}
}
