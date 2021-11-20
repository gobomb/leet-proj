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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
