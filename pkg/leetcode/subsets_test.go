package leetcode

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// func Test_subsets(t *testing.T) {
// 	type args struct {
// 		nums []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want [][]int
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				nums: []int{1, 2, 3},
// 			},
// 			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
// 		},
// 		{
// 			name: "2",
// 			args: args{
// 				nums: []int{0},
// 			},
// 			want: [][]int{{}, {0}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("subsets(%v) = %v, want %v", tt.args.nums, got, tt.want)
// 			}
// 		})
// 	}
// }
