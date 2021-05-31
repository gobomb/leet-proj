package leetcode

import (
	"container/list"
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_treePrinter_String(t *testing.T) {
	type fields struct {
		list *list.List
		root *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				list: listFree.Get().(*list.List).Init(),
				root: MakeTree([]int{1, 2, 3, 4, 5, 6}),
			},
			want: "",
		},
		{
			name: "2",
			fields: fields{
				list: listFree.Get().(*list.List).Init(),
				root: MakeTree([]int{1, Null, 3, Null, Null, 6}),
			},
			want: "",
		},
		{
			name: "3",
			fields: fields{
				list: listFree.Get().(*list.List).Init(),
				root: MakeTree([]int{1}),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &treePrinter{
				list: tt.fields.list,
				root: tt.fields.root,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("treePrinter.String() = %v, want %v", got, tt.want)
			}
		})
	}
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
