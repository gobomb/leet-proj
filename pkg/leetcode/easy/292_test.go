package easy

import "testing"

func Test_canWinNim(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				n: 4,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
