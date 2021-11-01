package leetcode

import (
	"reflect"
	"testing"
)

func Test_partitionPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				s: "aab",
			},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "2",
			args: args{
				s: "a",
			},
			want: [][]string{{"a"}},
		},
		{
			name: "3",
			args: args{
				s: "cbbbcc",
			},
			want: [][]string{{"c", "b", "b", "b", "c", "c"}, {"c", "b", "b", "b", "cc"}, {"c", "b", "bb", "c", "c"}, {"c", "b", "bb", "cc"}, {"c", "bb", "b", "c", "c"}, {"c", "bb", "b", "cc"}, {"c", "bbb", "c", "c"}, {"c", "bbb", "cc"}, {"cbbbc", "c"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionPalindrome(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
