package review

import "testing"

func Test_isValid11(t *testing.T) {
	type args struct {
		s string
	}

	ss := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: ss[0],
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: ss[1],
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: ss[2],
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: ss[3],
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: ss[4],
			},
			want: true,
		},

		{
			name: "",
			args: args{
				s: "]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid11() = %v, want %v", got, tt.want)
			}
		})
	}
}
