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
					makeListNode([]int{1, 4, 5}),
					makeListNode([]int{1, 3, 4}),
					makeListNode([]int{2, 6}),
				},
			},
			makeListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
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
					makeListNode([]int{}),
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
