package leetcode

import (
	"testing"
)

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				beginWord: "a",
				endWord:   "c",
				wordList:  []string{"a", "b", "c"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ladderLengthBiDBFS(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				beginWord: "a",
				endWord:   "c",
				wordList:  []string{"a", "b", "c"},
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				beginWord: "hot",
				endWord:   "dog",
				wordList:  []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLengthBiDBFS(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLengthBiDBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
