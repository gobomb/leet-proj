package medium

import "testing"

func Test_hIndexII(t *testing.T) {
	type args struct {
		citations []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				citations: []int{0, 1, 3, 5, 6},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				citations: []int{1, 2, 100},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				citations: []int{0},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				citations: []int{1},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				citations: []int{100},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndexII(tt.args.citations); got != tt.want {
				t.Errorf("hIndexII() = %v, want %v", got, tt.want)
			}
		})
	}
}
