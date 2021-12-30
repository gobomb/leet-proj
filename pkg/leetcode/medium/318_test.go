package medium

import (
	"reflect"
	"testing"
)

func Test_maxProduct1(t *testing.T) {
	type args struct {
		words []string
	}

	tests := func() []struct {
		name string
		args args
		want int
	} {
		return []struct {
			name string
			args args
			want int
		}{
			{
				name: "",
				args: args{
					words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
				},
				want: 16,
			},
			{
				name: "",
				args: args{
					words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
				},
				want: 4,
			},
			{
				name: "",
				args: args{
					words: []string{"a", "aa", "aaa", "aaaa"},
				},
				want: 0,
			},
		}
	}
	toTest := []func([]string) int{maxProduct1, maxProduct2}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.words); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
