package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_grayCode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				n: 2,
			},
			want: []int{0, 1, 3, 2},
		},
		{
			name: "2",
			args: args{
				n: 1,
			},
			want: []int{0, 1},
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
