package medium

import "testing"

func Test_getMoneyAmount(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 10,
			},
			want: 16,
		},
		{
			name: "",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.args.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
