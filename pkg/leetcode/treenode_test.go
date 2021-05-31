package leetcode

import (
	"container/list"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func init() {
	log.SetFlags(log.Lshortfile)
}

// func TestTreeStringer(t *testing.T) {
// 	type args struct {
// 		l *TreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want fmt.Stringer
// 	}{
// 		{
// 			name: "",
// 			args: args{
// 				l: MakeTree([]int{1, 2, 3, 4, 5, 6}),
// 			},
// 			want: nil,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := TreeStringer(tt.args.l); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TreeStringer(%v) = %v, want %v", tt.args.l, got, tt.want)
// 			}
// 		})
// 	}
// }

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
