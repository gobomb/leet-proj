package medium

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "1+3*2+2*3",
			},
			want: 13,
		},
		{
			name: "",
			args: args{
				s: "3+2*2",
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				s: " 3/2 ",
			},
			want: 1,
		},

		{
			name: "",
			args: args{
				s: " 3+5 / 2 ",
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
