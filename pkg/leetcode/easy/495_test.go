package easy

import "testing"

func Test_findPoisonedDuration(t *testing.T) {
	type args struct {
		timeSeries []int
		duration   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				timeSeries: []int{1, 4},
				duration:   2,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				timeSeries: []int{1, 2},
				duration:   2,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				timeSeries: []int{},
				duration:   2,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				timeSeries: []int{1},
				duration:   2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				timeSeries: []int{1, 2, 3, 5},
				duration:   2,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoisonedDuration(tt.args.timeSeries, tt.args.duration); got != tt.want {
				t.Errorf("findPoisonedDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
