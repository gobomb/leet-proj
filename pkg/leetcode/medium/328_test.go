package medium

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				MakeListNode(1, 2, 3, 4, 5),
			},
			want: MakeListNode(1, 3, 5, 2, 4),
		},
		{
			name: "",
			args: args{
				MakeListNode(2, 1, 3, 5, 6, 4, 7),
			},
			want: MakeListNode(2, 3, 6, 7, 1, 5, 4),
		},
		{
			name: "",
			args: args{
				MakeListNode(1, 2),
			},
			want: MakeListNode(1, 2),
		},
		{
			name: "",
			args: args{
				MakeListNode(1, 2, 3),
			},
			want: MakeListNode(1, 3, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
