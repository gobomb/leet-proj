package leetcode

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbbaccc",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s1: "",
				s2: "",
				s3: "",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s1: "c",
				s2: "ca",
				s3: "cac",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				s1: "",
				s2: "ca",
				s3: "",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				s1: "",
				s2: "ca",
				s3: "ca",
			},
			want: true,
		},
		{
			name: "7",
			args: args{
				s1: "a",
				s2: "c",
				s3: "ca",
			},
			want: true,
		},
		{
			name: "8",
			args: args{
				s1: "abababababababababababababababababababababababababababababababababababababababababababababababababbb",
				s2: "babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
				s3: "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInterleaveCache(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbbaccc",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s1: "",
				s2: "",
				s3: "",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s1: "c",
				s2: "ca",
				s3: "cac",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				s1: "",
				s2: "ca",
				s3: "",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				s1: "",
				s2: "ca",
				s3: "ca",
			},
			want: true,
		},
		{
			name: "7",
			args: args{
				s1: "a",
				s2: "c",
				s3: "ca",
			},
			want: true,
		},
		{
			name: "8",
			args: args{
				s1: "abababababababababababababababababababababababababababababababababababababababababababababababababbb",
				s2: "babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
				s3: "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleaveCache(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleaveCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInterleave2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbbaccc",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s1: "",
				s2: "",
				s3: "",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s1: "c",
				s2: "ca",
				s3: "cac",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				s1: "",
				s2: "ca",
				s3: "",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				s1: "",
				s2: "ca",
				s3: "ca",
			},
			want: true,
		},
		{
			name: "7",
			args: args{
				s1: "a",
				s2: "c",
				s3: "ca",
			},
			want: true,
		},
		// {
		// 	name: "8",
		// 	args: args{
		// 		s1: "abababababababababababababababababababababababababababababababababababababababababababababababababbb",
		// 		s2: "babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
		// 		s3: "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
		// 	},
		// 	want: false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave2(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleaveCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
