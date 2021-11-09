package easy

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
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
				head: MakeListNode(1, 2, 3, 4, 5),
			},
			want: MakeListNode(3, 4, 5),
		},
		{
			name: "0",
			args: args{
				head: MakeListNode(1, 2, 3, 4, 5, 6),
			},
			want: MakeListNode(4, 5, 6),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
