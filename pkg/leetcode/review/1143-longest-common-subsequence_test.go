package review

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		t1 string
		t2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				t1: "abcde",
				t2: "ace",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				t1: "abcde",
				t2: "acce",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				t1: "aceacce",
				t2: "acce",
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				t1: "abc",
				t2: "abc",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				t1: "abc",
				t2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
