package leetcode

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// func TestMakeTree2(t *testing.T) {
// 	ints := [][]int{[]int{1, Null, 2, Null, 3},
// 		[]int{1, Null, 3, 2},
// 		[]int{2, 1, 3},
// 		[]int{3, 1, Null, Null, 2}, []int{3, 2, Null, 1}}
// 	ints2 := [][]int{{1, Null, 2, Null, Null, Null, 3},
// 		{1, Null, 3, 2},
// 		{2, 1, 3},
// 		{3, 1, Null, Null, 2},
// 		{3, 2, Null, 1}}
// 	wants := []*TreeNode{}
// 	type args struct {
// 		vals []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *TreeNode
// 	}{}
// 	for i := range ints {
// 		wants = append(wants, MakeTree(ints2[i]))
// 		tests = append(tests, struct {
// 			name string
// 			args args
// 			want *TreeNode
// 		}{
// 			"",
// 			args{ints[i]},
// 			wants[i],
// 		})
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := MakeTree2(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("MakeTree2(%+v) = %v, want %+v", tt.args.vals, got, tt.want)
// 			} else {
// 				// log.Printf("%+v\n", got)
// 			}
// 		})
// 	}
// }
