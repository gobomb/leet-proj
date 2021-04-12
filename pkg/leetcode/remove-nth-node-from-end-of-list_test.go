package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: makeListNode([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: makeListNode([]int{1, 2, 3, 5}),
		},
		{
			name: "2",
			args: args{
				head: makeListNode([]int{1}),
				n:    1,
			},
			want: makeListNode([]int{}),
		},
		{
			name: "3",
			args: args{
				head: makeListNode([]int{1, 2}),
				n:    1,
			},
			want: makeListNode([]int{1}),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
