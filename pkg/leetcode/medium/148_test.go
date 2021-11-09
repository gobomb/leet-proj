package medium

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	genTests := func() []struct {
		name string
		args args
		want *ListNode
	} {
		return []struct {
			name string
			args args
			want *ListNode
		}{
			{
				name: "0",
				args: args{
					head: MakeListNode(6, 5, 3, 1, 8, 7, 2, 4),
				},
				want: MakeListNode(1, 2, 3, 4, 5, 6, 7, 8),
			},
			{
				name: "1",
				args: args{
					head: MakeListNode(4, 2, 1, 3),
				},
				want: MakeListNode(1, 2, 3, 4),
			},
			{
				name: "3",
				args: args{
					head: MakeListNode(-1, 5, 3, 4, 0),
				},
				want: MakeListNode(-1, 0, 3, 4, 5),
			},
		}
	}

	toTest := []func(*ListNode) *ListNode{sortList, sortListMerge, sortList1}

	for _, f := range toTest {
		for _, tt := range genTests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.head); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v() = %+v, want %+v", funcName(f), got, tt.want)
				}
			})
		}
	}
}
