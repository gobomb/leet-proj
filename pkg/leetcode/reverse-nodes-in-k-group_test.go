package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{

		{
			"1",
			args{
				makeListNode([]int{1, 2, 3, 4, 5}),
				2,
			},
			makeListNode([]int{2, 1, 4, 3, 5}),
		},
		{
			"2",
			args{
				makeListNode([]int{1, 2, 3, 4, 5}),
				3,
			},
			makeListNode([]int{3, 2, 1, 4, 5}),
		},
		{
			"3",
			args{
				makeListNode([]int{1, 2, 3, 4, 5}),
				1,
			},
			makeListNode([]int{1, 2, 3, 4, 5}),
		},
		{
			"4",
			args{
				makeListNode([]int{1}),
				3,
			},
			makeListNode([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		first *ListNode
		last  *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				makeListNode([]int{1, 2, 3, 4, 5}),
				nil,
			},
			makeListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			"1",
			args{
				makeListNode([]int{1, 2}),
				nil,
			},
			makeListNode([]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.first, tt.args.last); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
