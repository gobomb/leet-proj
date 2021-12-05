package medium

import "testing"

func Test_computeArea(t *testing.T) {
	type args struct {
		ax1 int
		ay1 int
		ax2 int
		ay2 int
		bx1 int
		by1 int
		bx2 int
		by2 int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				ax1: -3,
				ay1: 0,
				ax2: 3,
				ay2: 4,
				bx1: 0,
				by1: -1,
				bx2: 9,
				by2: 2,
			},
			want: 45,
		},
		{
			name: "",
			args: args{
				ax1: -2,
				ay1: -2,
				ax2: 2,
				ay2: 2,
				bx1: 3,
				by1: 3,
				bx2: 4,
				by2: 4,
			},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeArea(tt.args.ax1, tt.args.ay1, tt.args.ax2, tt.args.ay2, tt.args.bx1, tt.args.by1, tt.args.bx2, tt.args.by2); got != tt.want {
				t.Errorf("computeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
