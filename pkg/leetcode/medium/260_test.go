package medium

import (
	"reflect"
	"testing"
)

func Test_singleNumberIII(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 1, 3, 2, 5},
			},
			want: []int{3, 5},
		},
		{
			name: "",
			args: args{
				nums: []int{-1, 0},
			},
			want: []int{-1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberIII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumberIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
