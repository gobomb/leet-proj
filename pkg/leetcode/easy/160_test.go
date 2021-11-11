package easy

import (
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	ha1 := MakeListNode(4, 1)
	hb1 := MakeListNode(5, 6, 1)
	inter1 := MakeListNode(8, 4, 5)
	ha1.Next.Next = inter1
	hb1.Next.Next.Next = inter1

	ha2 := MakeListNode(2, 6, 4)
	hb2 := MakeListNode(1, 5)

	ha3 := MakeListNode(4, 1, 8, 4, 5)
	hb3 := MakeListNode(5, 6, 1, 8, 4, 5)

	// ha4 := MakeListNode(2, 6, 4)
	// hb4 := MakeListNode(1, 5)
	// inter4 := MakeListNode(0)
	// ha4.Next.Next = inter4
	// hb4.Next.Next.Next = inter4

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				headA: ha1,
				headB: hb1,
			},
			want: inter1,
		},
		{
			name: "2",
			args: args{
				headA: ha2,
				headB: hb2,
			},
			want: nil,
		},
		{
			name: "3",
			args: args{
				headA: ha3,
				headB: hb3,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); got != tt.want {
				t.Errorf("getIntersectionNode() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
