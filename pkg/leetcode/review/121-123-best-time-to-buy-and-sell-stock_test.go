package review

import "testing"

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				prices: []int{7},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitIII(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				prices: []int{7},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				prices: []int{6, 1, 3, 2, 4, 7},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIII(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitIIIdp(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				prices: []int{7},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				prices: []int{6, 1, 3, 2, 4, 7},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIIIdp(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIIIdp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitIIIdp2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				prices: []int{7},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				prices: []int{6, 1, 3, 2, 4, 7},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIIIdp2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIIIdp2() = %v, want %v", got, tt.want)
			}
		})
	}
}
