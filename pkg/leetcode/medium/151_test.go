package medium

import (
	"reflect"
	"testing"
)

func Test_reverseWords(t *testing.T) {
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
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		},
		{
			name: "",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
		{
			name: "",
			args: args{
				s: "F R  I   E    N     D      S      ",
			},
			want: "S D N E I R F",
		},
	}

	toTest := [](func(s string) string){reverseWords, reverseWordsWithOutStl}

	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.s); got != tt.want {
					t.Errorf("%v() = %v, want %v", funcName(f), got, tt.want)
				}
			})
		}
	}
}

func Test_trimSpace(t *testing.T) {
	type args struct {
		b []byte
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "去除开头的空格",
			args: args{
				b: []byte(string("    aabb")),
			},
			want: []byte(string("aabb")),
		},
		{
			name: "去除中间的空格",
			args: args{
				b: []byte(string("   aabb    def    g      glklsd")),
			},
			want: []byte(string("aabb def g glklsd")),
		},
		{
			name: "去除结尾的空格",
			args: args{
				b: []byte(string("   aabb    def    g      glklsd      ")),
			},
			want: []byte(string("aabb def g glklsd")),
		},
		{
			name: "",
			args: args{
				b: []byte(string("F R  I   E    N     D      S      ")),
			},
			want: []byte("F R I E N D S"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimSpace(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimSpace() = %s, want %s", got, tt.want)
			}
		})
	}
}
