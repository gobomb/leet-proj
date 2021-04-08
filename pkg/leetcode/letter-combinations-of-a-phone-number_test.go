package leetcode

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{

			"1",
			args{"23"},
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{

			"2",
			args{"2"},
			[]string{"a", "b", "c"},
		},
		{

			"3",
			args{""},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_find(t *testing.T) {
	type args struct {
		index  int
		digits string
		s      string
	}
	tests := []struct {
		name   string
		args   args
		wantRs []string
	}{
		{
			"1",
			args{1, "23", "a"},
			[]string{"ad", "ae", "af"},
		},
		{
			"2",
			args{1, "2", "a"},
			[]string{"a"},
		},
		{
			"3",
			args{0, "23", ""},
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := find(tt.args.index, tt.args.digits, tt.args.s); !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("find() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}
