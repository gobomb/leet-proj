package medium

import (
	"fmt"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
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
				nums: []int{1, 5, 1, 1, 6, 4},
			},
			want: []int{1, 5, 1, 6, 1, 4},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 3, 2, 2, 3, 1},
			},
			want: []int{2, 3, 1, 3, 1, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 5, 2, 2, 1, 4},
			},
			want: []int{2, 4, 1, 5, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := fmt.Sprintf("%v", tt.want)
			wiggleSort(tt.args.nums)
			got := fmt.Sprintf("%v", tt.args.nums)
			if want != got {
				t.Fatalf("failed! want %s got %s ", want, got)
			}
		})
	}
}
