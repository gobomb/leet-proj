package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_deleteDuplicates(t *testing.T) {
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
				head: makeListNode([]int{1, 1, 2}),
			},
			want: makeListNode([]int{1, 2}),
		},
		{
			name: "2",
			args: args{
				head: makeListNode([]int{1, 1, 2, 3, 3}),
			},
			want: makeListNode([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates(%+v) = %+v, want %+v", tt.args.head, got, tt.want)
			}
		})
	}
}
