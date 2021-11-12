package easy

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: MakeListNode(1, 2, 6, 3, 4, 5, 6),
				val:  6,
			},
			want: MakeListNode(1, 2, 3, 4, 5),
		},
		{
			name: "",
			args: args{
				head: MakeListNode(),
				val:  1,
			},
			want: MakeListNode(),
		},
		{
			name: "",
			args: args{
				head: MakeListNode(7, 7, 7, 7),
				val:  7,
			},
			want: MakeListNode(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
