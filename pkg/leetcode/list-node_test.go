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

	fmt.Printf("loop %+v\n", loop)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := LNDeepCopy(tt.args.o); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("LNDeepCopy() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"1make",
			fields{
				Val:  0,
				Next: makeListNode([]int{1, 2, 4}),
			},
			"0->1->2->4->(nil)",
		},
		{
			"2",
			fields{
				Val:  0,
				Next: nil,
			},
			"0->(nil)",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("ListNode.String() = %v, want %v", got, tt.want)
			}
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
	want := "1->3->(loop)"
	if got := loop.String(); got != want {
		t.Errorf("ListNode.String() = %v, want %v", got, want)
	}
}
