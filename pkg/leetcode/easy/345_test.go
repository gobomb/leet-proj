package easy

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "hello",
			},
			want: "holle",
		},

		{
			name: "",
			args: args{
				s: "leetcode",
			},
			want: "leotcede",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
