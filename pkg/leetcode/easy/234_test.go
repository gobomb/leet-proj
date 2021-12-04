package easy

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				head: MakeListNode(1, 2, 2, 1),
			},
			want: true,
		}, {
			name: "",
			args: args{
				head: MakeListNode(1, 2, 3, 4, 3, 2, 1),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				head: MakeListNode(1, 2),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				head: MakeListNode(1, 2, 2, 1, 3, 4, 4),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
