package easy

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				nums: []int{6, 6, 6, 7, 7},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
