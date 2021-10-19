package leetcode

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				words: []string{"Hello", "Alaska", "Dad", "Peace"},
			},
			want: []string{"Alaska", "Dad"},
		},
		{
			name: "2",
			args: args{
				words: []string{"omk"},
			},
			want: []string{},
		},
		{
			name: "3",
			args: args{
				words: []string{"adsdf", "sfd"},
			},
			want: []string{"adsdf", "sfd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
