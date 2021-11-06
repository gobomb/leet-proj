package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				lists: []*ListNode{
					MakeListNode(1, 4, 5),
					MakeListNode(1, 3, 4),
					MakeListNode(2, 6),
				},
			},
			MakeListNode(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			"2",
			args{
				[]*ListNode{},
			},
			nil,
		},
		{
			"3",
			args{
				[]*ListNode{
					MakeListNode(),
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
