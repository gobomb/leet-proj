package leetcode

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_reverseList(t *testing.T) {
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
				head: MakeListNode(1, 2, 3, 4, 5),
			},
			want: MakeListNode(5, 4, 3, 2, 1),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(1, 2),
			},
			want: MakeListNode(2, 1),
		},
		{
			name: "3",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argString := fmt.Sprintf("%+v", tt.args.head)
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList(%+v) = %+v, want %+v", argString, got, tt.want)
			}
		})
	}
}
