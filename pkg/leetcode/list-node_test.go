package leetcode

import (
	"reflect"
	"testing"
)

func Test_makeListNode(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"124",
			args{
				[]int{1, 2, 4},
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						4,
						nil,
					},
				},
			},
		},
		{
			"112244",
			args{
				[]int{1, 1, 2, 2, 4, 4},
			},
			&ListNode{
				1,
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							2,
							&ListNode{
								4,
								&ListNode{
									4,
									nil,
								},
							},
						},
					},
				},
			},
		},
		{
			"nil",
			args{
				[]int{},
			},
			nil,
		},
		{
			"0",
			args{[]int{0}},
			&ListNode{
				0,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeListNode(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
