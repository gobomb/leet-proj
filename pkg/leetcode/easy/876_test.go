package easy

import (
	"log"
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
			name: "1",
			args: args{
				head: MakeListNode(1, 2, 3, 4, 5, 6),
			},
			want: MakeListNode(4, 5, 6),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(6, 5, 3, 1, 8, 7, 2, 4),
			},
			want: MakeListNode(8, 7, 2, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			} else {
				log.Printf("%+v", got)
			}
		})
	}
}
