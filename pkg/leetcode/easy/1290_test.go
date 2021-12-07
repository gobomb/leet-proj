package easy

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				head: MakeListNode(1, 0, 1),
			},
			want: 5,
		}, {
			name: "",
			args: args{
				head: MakeListNode(0),
			},
			want: 0,
		}, {
			name: "",
			args: args{
				head: MakeListNode(1),
			},
			want: 1,
		}, {
			name: "",
			args: args{
				head: MakeListNode(1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0),
			},
			want: 18880,
		}, {
			name: "",
			args: args{
				head: MakeListNode(0, 0),
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
