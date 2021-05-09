package leetcode

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		// want [][]string
		want int
	}{
		{
			"1",
			args{
				4,
			},
			2,
			// [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			"2",
			args{
				1,
			},
			1,
		},
		{
			"3",
			args{
				3,
			},
			0,
		},
		{
			"5",
			args{
				5,
			},
			10,
		},
		{
			"6",
			args{
				6,
			},
			4,
		},
		{
			"7",
			args{
				7,
			},
			40,
		},
		{
			"8",
			args{
				8,
			},
			92,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("solveNQueens() = %v len= %v, want %v", got, len(got), tt.want)
			}
		})
	}
}
