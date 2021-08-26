package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head:  MakeListNode(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
			want: MakeListNode(1, 4, 3, 2, 5),
		},
		{
			name: "2",
			args: args{
				head:  MakeListNode(5),
				left:  1,
				right: 1,
			},
			want: MakeListNode(5),
		},
		{
			name: "3",
			args: args{
				head:  MakeListNode(3, 5),
				left:  1,
				right: 2,
			},
			want: MakeListNode(5, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
