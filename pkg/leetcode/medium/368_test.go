package medium

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
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
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 4, 8},
			},
			want: []int{1, 2, 4, 8},
		},
		{
			name: "",
			args: args{
				nums: []int{9, 18, 54, 90, 108, 180, 360, 540, 720},
			},
			want: []int{9, 18, 90, 180, 360, 720},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
