package tree

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var trees = [][]int{{1, Null, 2, Null, Null, Null, 3},
	{2, 1, 3},
	{3, 1, Null, Null, 2},
	{3, 2, Null, 1},
	{1, 2, Null, 3, 4, Null, Null, Null, 5, 6},
}

func TestMakeTree(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				vals: []int{1, 2, 3},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			name: "2",
			args: args{
				vals: []int{1, Null, 2, Null, Null, 3},
			},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTree(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTree(%v) = %v, want %v", tt.args.vals, got, tt.want)
			} else {
				log.Printf("%+v\n", got)
			}
		})
	}
}
func Test_printer(t *testing.T) {
	t.Run("test print nil tree", func(t *testing.T) {
		v := (*TreeNode)(nil)
		w := "Null"
		if got := TreeStringer(v).String(); got != w {
			t.Errorf("print treenode got %v ,want %v", got, w)
		}
		t.Logf("%+v", v)
	})
	t.Run("test print one node tree ", func(t *testing.T) {
		v := &TreeNode{}
		w := "0"
		if got := TreeStringer(v).String(); got != w {
			t.Errorf("print treenode got %v ,want %v", got, w)
		}
		t.Logf("%+v", v)
	})
	t.Run("test print node tree height", func(t *testing.T) {
		v := MakeTree(trees[4])
		w := "4"
		if got := fmt.Sprintf("%h", v); got != w {
			t.Errorf("print treenode got %v ,want %v", got, w)
		}
		t.Logf("%h", v)
	})
	t.Run("test print v/fallthrough", func(t *testing.T) {
		v := MakeTree(trees[4])
		w := "&TreeNode:{ 1 &TreeNode:{ 2 &TreeNode:{ 3 <nil> &TreeNode:{ 5 <nil> <nil> } } &TreeNode:{ 4 &TreeNode:{ 6 <nil> <nil> } <nil> } } <nil> }"
		if got := fmt.Sprintf("%v", v); got != w {
			t.Errorf("print treenode got %v ,want %v", got, w)
		}
		t.Logf("%v", v)
	})
	// t.Run("test print pointer", func(t *testing.T) {
	// 	var v *TreeNode
	// 	v = MakeTree(trees[4])
	// 	w := fmt.Sprintf("%p", unsafe.Pointer(v))
	// 	if got := fmt.Sprintf("%p", v); got != w {
	// 		t.Errorf("print treenode got %v ,want %v", got, w)
	// 	}
	// 	t.Logf("%v", w)
	// })
}
func Test_getTreeHeight(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: MakeTree(trees[0]),
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				root: MakeTree(trees[1]),
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				root: MakeTree(trees[2]),
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				root: MakeTree(trees[3]),
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				root: MakeTree(trees[4]),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTreeHeight(tt.args.root); got != tt.want {
				t.Errorf("getTreeHeight(%v) = %v, want %v", tt.args.root, got, tt.want)
			} else {
				log.Printf("%+v\n", tt.args.root)
			}
		})
	}
}

func TestTreeNode_Insert(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields *TreeNode
		args   args
		want   *TreeNode
	}{
		{
			"1",
			MakeTree([]int{1}),
			args{2},
			MakeTree([]int{1, Null, 2}),
		},
		{
			"2",
			MakeTree([]int{1, Null, 4}),
			args{3},
			MakeTree([]int{1, Null, 4, Null, Null, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Insert(tt.args.v)
			v1 := fmt.Sprintf("%+v", tt.fields)
			v2 := fmt.Sprintf("%+v", tt.want)
			if v1 != v2 {
				t.Errorf("failed insert %v got: %v \n want: %v", tt.args.v, v1, v2)
			}
		})
	}
}
