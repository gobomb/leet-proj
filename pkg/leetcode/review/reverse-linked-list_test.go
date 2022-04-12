package review

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
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
				name: "1",
				args: args{
					head: MakeListNode(1, 2, 3, 4, 5),
				},
				want: MakeListNode(5, 4, 3, 2, 1),
			},
			{
				name: "2",
				args: args{
					head: MakeListNode(1, 2),
				},
				want: MakeListNode(2, 1),
			},
			{
				name: "3",
				args: args{
					head: nil,
				},
				want: nil,
			},
		}
	}
	toTest := []func(*ListNode) *ListNode{reverseList2, reverseList}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.head); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
