package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "2",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "3",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "4",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
		{
			name: "5",
			args: args{
				path: "/a/../../b/../c//.//",
			},
			want: "/c",
		},
		{
			name: "6",
			args: args{
				path: "/a//b////c/d//././/..",
			},
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath(%v) = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath2(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath2(%v) = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
}
