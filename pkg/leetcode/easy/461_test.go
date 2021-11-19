package easy

import "testing"

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				x: 1,
				y: 4,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				x: 3,
				y: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				x: 28,
				y: 7,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
