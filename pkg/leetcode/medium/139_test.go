package medium

import (
	"reflect"
	"sort"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordBreakII(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{"cats and dog", "cat sand dog"},
		},
		{
			name: "2",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			name: "3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakII(tt.args.s, tt.args.wordDict)
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreakII() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
