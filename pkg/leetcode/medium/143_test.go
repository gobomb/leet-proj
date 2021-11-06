package medium

import (
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
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
				head: MakeListNode(1, 2, 3, 4, 5, 6),
			},
			want: MakeListNode(1, 6, 2, 5, 3, 4),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(1, 2, 3, 4, 5),
			},
			want: MakeListNode(1, 5, 2, 4, 3),
		},
		{
			name: "0 node",
			args: args{
				head: MakeListNode(),
			},
			want: MakeListNode(),
		},
		{
			name: "1 node",
			args: args{
				head: MakeListNode(1),
			},
			want: MakeListNode(1),
		},
		{
			name: "2 node",
			args: args{
				head: MakeListNode(1, 2),
			},
			want: MakeListNode(1, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if !reflect.DeepEqual(tt.args.head, tt.want) {
				t.Fatalf("output %+v want %+v \n", tt.args.head, tt.want)
			}
		})
	}
}
