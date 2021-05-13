package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
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
			makeListNode([]int{4, 5, 1, 2, 3}),
		},
		{
			"2",
			args{
				makeListNode([]int{0, 1, 2}),
				4,
			},
			makeListNode([]int{2, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
