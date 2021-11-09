package easy

import "testing"

func Test_truncateSentence(t *testing.T) {
	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "Hello how are you Contestant",
				k: 4,
			},
			want: "Hello how are you",
		},
		{
			name: "2",
			args: args{
				s: "What is the solution to this problem",
				k: 4,
			},
			want: "What is the solution",
		},
		{
			name: "3",
			args: args{
				s: "chopper is not a tanuki",
				k: 5,
			},
			want: "chopper is not a tanuki",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateSentence(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("truncateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
