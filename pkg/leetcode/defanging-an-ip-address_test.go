package leetcode

import "testing"

func Test_defangIPaddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				address: "1.1.1.1",
			},
			want: "1[.]1[.]1[.]1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defangIPaddr(tt.args.address); got != tt.want {
				t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
