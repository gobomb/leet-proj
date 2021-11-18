package medium

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				nums: []int{10, 2},
			},
			want: "210",
		},
		{
			name: "",
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "9534330",
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: "1",
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0},
			},
			want: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
