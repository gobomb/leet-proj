package leetcode

import "testing"

// ref https://jaycechant.info/2020/rapidly-generate-unit-tests-in-vs-code/
func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "1",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "2",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "3",
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			name: "4",
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			name: "5",
			s:    "mississippi",
			p:    "mis*is*p*.",
			want: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
