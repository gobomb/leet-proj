package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type matrix [][]byte

func (m matrix) String() string {
	var r string

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			r += fmt.Sprintf("%v ", string((m)[i][j]))
			if j == len(m)-1 {
				r += "\n" //fmt.Sprintf("%s\n", r)
			}
		}
	}
	// fmt.Println(m[1][1])

	return r
}

type argsSudoku struct {
	board [][]byte
}

var testsSudoku = []struct {
	name string
	args argsSudoku
	want [][]byte
}{
	{
		"1",
		argsSudoku{
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
		[][]byte{
			{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
			{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
			{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
			{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
			{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
			{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
			{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
			{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
			{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
		},
	},
}

func Test_solveSudoku(t *testing.T) {
	for _, tt := range testsSudoku {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Fatalf("solveSudoku return \n%v \n want \n%v \n", matrix(tt.args.board), matrix(tt.want))
			}
		})
	}
}

func Test_checkSudoku(t *testing.T) {
	type args struct {
		board *[][]byte
		pos   position
		val   byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				&[][]byte{
					{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
				position{2, 0},
				'1',
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSudoku(tt.args.board, tt.args.pos, tt.args.val); got != tt.want {
				t.Errorf("checkSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
