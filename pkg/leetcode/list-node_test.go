package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_makeListNode(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"124",
			args{
				[]int{1, 2, 4},
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						4,
						nil,
					},
				},
			},
		},
		{
			"112244",
			args{
				[]int{1, 1, 2, 2, 4, 4},
			},
			&ListNode{
				1,
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							2,
							&ListNode{
								4,
								&ListNode{
									4,
									nil,
								},
							},
						},
					},
				},
			},
		},
		{
			"nil",
			args{
				[]int{},
			},
			nil,
		},
		{
			"0",
			args{[]int{0}},
			&ListNode{
				0,
				nil,
			},
		},
		{
			"1,2",
			args{[]int{1, 2}},
			&ListNode{
				1,
				&ListNode{
					2,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeListNode(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLNDeepCopy(t *testing.T) {
	type args struct {
		o *ListNode
	}
	loop := &ListNode{}
	loop.Val = 1
	loop.Next = &ListNode{
		3,
		loop,
	}

	tests := []struct {
		name  string
		args  args
		wantN *ListNode
	}{
		// TODO: Add test cases.
		{
			"1make",
			args{
				makeListNode([]int{1, 2, 4}),
			},
			makeListNode([]int{1, 2, 4}),
		},
		{
			"2",
			args{
				nil,
			},
			nil,
		},
		{
			"3",
			args{
				&ListNode{
					0,
					nil,
				},
			},
			&ListNode{
				0,
				nil,
			},
		},
		{
			"loopCheck",
			args{
				loop,
			},
			loop,
		},
	}

	// fmt.Printf("loop %v\n", NewListPrint(loop))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := LNDeepCopy(tt.args.o); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("LNDeepCopy() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

type fieldsString struct {
	Val  int
	Next *ListNode
}

var testsString = []struct {
	name   string
	fields fieldsString
	want   string
}{
	{
		"1make",
		fieldsString{
			Val:  0,
			Next: makeListNode([]int{1, 2, 4}),
		},
		"0->1->2->4->(nil)",
	},
	{
		"2",
		fieldsString{
			Val:  0,
			Next: nil,
		},
		"0->(nil)",
	},
	// TODO: Add test cases.
}

func TestListNode_Format(t *testing.T) {
	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := fmt.Sprintf("%+v", l); got != tt.want {
				t.Errorf("ListNode.String() = %v, want %v", got, tt.want)
			}
			fmt.Printf("+v %+v\t", l)
			fmt.Printf("v %v\n", l)
		})
	}
}
func TestListNode_String_loop(t *testing.T) {
	loop := &ListNode{}
	loop.Val = 1
	loop.Next = &ListNode{
		3,
		loop,
	}
	want := "1->3->(loop)1"
	if got := ListStringer(loop).String(); got != want {
		t.Errorf("ListNode.String() = %v, want %v", got, want)
	}
	t.Logf("%v\n", ListStringer(loop))

}

func TestListPrinter_String(t *testing.T) {
	ln := makeListNode([]int{1, 2, 4})
	type fields struct {
		visitMap visitMap
		head     *ListNode
		cur      *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"1",
			fields{
				visitMap: make(visitMap),
				cur:      ln,
				head:     ln, //  makeListNode([]int{1, 2, 4}),
			},
			"1->2->4->(nil)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListPrinter{
				visitMap: tt.fields.visitMap,
				head:     tt.fields.head,
				cur:      tt.fields.cur,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("ListPrinter.String() = %v, want %v", got, tt.want)
			}
			t.Logf("%v\n", l)
		})
	}

	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := ListStringer(l).String(); got != tt.want {
				t.Errorf("ListNode.String() = %v, want %v", got, tt.want)
			}
			t.Logf("%v\n", ListStringer(l))
		})
	}
}
