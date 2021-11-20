package medium

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				s: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			},
			want: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name: "",
			args: args{
				s: "AAAAAAAAAAAAA",
			},
			want: []string{"AAAAAAAAAA"},
		},
	}

	toTest := []func(string) []string{findRepeatedDnaSequences, findRepeatedDnaSequences1}

	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v() = %v, want %v",funcName(f), got, tt.want)
				}
			})
		}
	}
}
