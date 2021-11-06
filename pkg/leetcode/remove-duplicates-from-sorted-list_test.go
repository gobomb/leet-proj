package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_deleteDuplicates(t *testing.T) {
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
				head: MakeListNode(1, 1, 2),
			},
			want: MakeListNode(1, 2),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(1, 1, 2, 3, 3),
			},
			want: MakeListNode(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates(%+v) = %+v, want %+v", tt.args.head, got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicatesII(t *testing.T) {
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
				head: MakeListNode(1, 1, 2),
			},
			want: MakeListNode(1, 2),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(1, 1, 2, 3, 3),
			},
			want: MakeListNode(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesII(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicatesII() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicatesII2(t *testing.T) {
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
				head: MakeListNode(1, 1, 2),
			},
			want: MakeListNode(1, 2),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(1, 1, 2, 3, 3),
			},
			want: MakeListNode(1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesII2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicatesII2() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
