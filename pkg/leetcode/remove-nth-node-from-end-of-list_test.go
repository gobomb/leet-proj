package leetcode

import (
	"reflect"
	"testing"
)

type argsRmNthFE struct {
	head *ListNode
	n    int
}

var testsRmNthFE = []struct {
	name string
	args argsRmNthFE
	want *ListNode
}{
	{
		name: "1",
		args: argsRmNthFE{
			head: makeListNode([]int{1, 2, 3, 4, 5}),
			n:    2,
		},
		want: makeListNode([]int{1, 2, 3, 5}),
	},
	{
		name: "2",
		args: argsRmNthFE{
			head: makeListNode([]int{1}),
			n:    1,
		},
		want: makeListNode([]int{}),
	},
	{
		name: "3",
		args: argsRmNthFE{
			head: makeListNode([]int{1, 2}),
			n:    1,
		},
		want: makeListNode([]int{1}),
	},
}

func Test_removeNthFromEnd(t *testing.T) {
	for _, tt := range testsRmNthFE {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(LNDeepCopy(tt.args.head), tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNthFromEndRangeTwo(t *testing.T) {
	for _, tt := range testsRmNthFE {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEndRangeTwo(LNDeepCopy(tt.args.head), tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEndRangeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNthFromEndFastSlowP(t *testing.T) {
	for _, tt := range testsRmNthFE {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEndFastSlowP(LNDeepCopy(tt.args.head), tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEndFastSlowP() = %v, want %v", got, tt.want)
			}
		})
	}
}
