package review

import (
	"reflect"
	"testing"
)

// 25. Reverse Nodes in k-Group

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
				MakeListNode(1, 2, 3, 4, 5),
				2,
			},
			MakeListNode(2, 1, 4, 3, 5),
		},
		{
			"2",
			args{
				MakeListNode(1, 2, 3, 4, 5),
				3,
			},
			MakeListNode(3, 2, 1, 4, 5),
		},
		{
			"3",
			args{
				MakeListNode(1, 2, 3, 4, 5),
				1,
			},
			MakeListNode(1, 2, 3, 4, 5),
		},
		{
			"4",
			args{
				MakeListNode(1),
				3,
			},
			MakeListNode(1),
		},
		{
			"5",
			args{
				MakeListNode(1, 2),
				2,
			},
			MakeListNode(2, 1),
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