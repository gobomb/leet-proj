package leetcode

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: MakeListNode(1, 4, 3, 2, 5, 2),
				x:    3,
			},
			want: MakeListNode(1, 2, 2, 4, 3, 5),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(2, 1),
				x:    2,
			},
			want: MakeListNode(1, 2),
		},
		{
			name: "3",
			args: args{
				head: MakeListNode(2),
				x:    2,
			},
			want: MakeListNode(2),
		},
		{
			name: "4",
			args: args{
				head: MakeListNode(2, 1, 2, 1, 3, 4, 1, 5, 7),
				x:    4,
			},
			want: MakeListNode(2, 1, 2, 1, 3, 1, 4, 5, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
