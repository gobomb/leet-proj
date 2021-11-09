package medium

import (
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "0",
			args: args{
				head: MakeListNode(6, 5, 3, 1, 8, 7, 2, 4),
			},
			want: MakeListNode(1, 2, 3, 4, 5, 6, 7, 8),
		},
		{
			name: "1",
			args: args{
				head: MakeListNode(4, 2, 1, 3),
			},
			want: MakeListNode(1, 2, 3, 4),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(-1, 5, 3, 4, 0),
			},
			want: MakeListNode(-1, 0, 3, 4, 5),
		},
		{
			name: "3",
			args: args{
				head: MakeListNode(-1),
			},
			want: MakeListNode(-1),
		},
		{
			name: "4",
			args: args{
				head: MakeListNode(-1, 0, 0, 0, -3),
			},
			want: MakeListNode(-3, -1, 0, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
