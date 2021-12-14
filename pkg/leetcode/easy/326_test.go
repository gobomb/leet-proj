package easy

import "testing"

func Test_isPowerOfThree(t *testing.T) {
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
				n: 27,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n: 9,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 45,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
