package easy

import (
	"reflect"
	"testing"
)

/*
	2. Add Two Numbers
*/

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := func() []struct {
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
				name: "",
				args: args{
					l1: MakeListNode(2, 4, 3),
					l2: MakeListNode(5, 6, 4),
				},
				want: MakeListNode(7, 0, 8),
			},
			{
				name: "",
				args: args{
					l1: MakeListNode(0),
					l2: MakeListNode(0),
				},
				want: MakeListNode(0),
			},
			{
				name: "",
				args: args{
					l1: MakeListNode(9, 9, 9),
					l2: MakeListNode(1, 1, 1),
				},
				want: MakeListNode(0, 1, 1, 1),
			},
		}
	}
	
	toTest := []func(l1, l2 *ListNode) *ListNode{addTwoNumbers}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
