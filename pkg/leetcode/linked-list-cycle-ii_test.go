package leetcode

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: MakeListNode(-21, 10, 17, 8, 4,
					26, 15, 35, 33, -7,
					-16, 27, -12, 6, 29,
					-12, 15, 9, 20, 14,
					14, 2, 13, -24, 21,
					23, -21, 5),
			},
		},
	}
	tail := tests[0].args.head.FindNode(21)
	tests[0].args.head.FindNode(5).Next = tail
	tests[0].want = tail
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
