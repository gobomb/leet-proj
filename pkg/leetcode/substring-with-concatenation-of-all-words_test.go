package leetcode

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				"barfoothefoobarman",
				[]string{"foo", "bar"},
			},
			[]int{0, 9},
		},
		{
			"2",
			args{
				"barfoofoobarthefoobarman",
				[]string{"bar", "foo", "the"},
			},
			[]int{6, 9, 12},
		},
		{
			"3",
			args{
				"wordgoodgoodgoodbestword",
				[]string{"word", "good", "best", "word"},
			},
			[]int{},
		},
		{
			"4",
			args{
				"wordgoodgoodgoodbestword",
				[]string{"word", "good", "best", "good"},
			},
			[]int{8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_matchWord(t *testing.T) {
	type args struct {
		i    int
		s    string
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				3,
				"barfoothefoobarman",
				"foo",
			},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchWord(tt.args.i, tt.args.s, tt.args.word); got != tt.want {
				t.Errorf("matchWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
