package easy

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candyType []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				candyType: []int{1, 1, 2, 2, 3, 3},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				candyType: []int{1, 1, 2, 3},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				candyType: []int{6, 6, 6, 6},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candyType); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
