package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{

		{
			name: "1",
			args: args{
				s: []byte{'a', 'b', 'c'},
			},
			want: []byte{'c', 'b', 'a'},
		},
		{
			name: "2",
			args: args{
				s: []byte{'a', 'b', 'c', 'd'},
			},
			want: []byte{'d', 'c', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temp := make([]byte, len(tt.args.s))
			copy(temp, tt.args.s)
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.want, tt.args.s) {
				t.Errorf("failed want %v ,args %v got %v", tt.want, temp, tt.args.s)
			}
		})
	}
}
