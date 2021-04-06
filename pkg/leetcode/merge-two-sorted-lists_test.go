package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				l1: &ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							4,
							nil,
						},
					},
				},

				l2: &ListNode{
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
			"2",
			args{
				l1: nil,
				l2: nil,
			},
			nil,
		},
		{
			"3",
			args{
				l1: nil,
				l2: &ListNode{
					0,
					nil,
				},
			},
			&ListNode{
				0,
				nil,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
