package leetcode

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		res  [][]byte
	}{
		{
			name: "1",
			args: args{
				board: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			},
			res: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			name: "2",
			args: args{
				board: [][]byte{{'X'}},
			},
			res: [][]byte{{'X'}},
		},
		{
			name: "3",
			args: args{
				board: [][]byte{{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'}, {'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'O'}, {'O', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'X', 'O', 'X', 'X', 'O', 'O', 'X', 'X', 'X'}, {'O', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'}, {'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O'}, {'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'}},
			},
			res: [][]byte{{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'}, {'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'}, {'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}, {'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'}, {'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O', 'O'}, {'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			for i := range tt.args.board {
				for j := range tt.args.board[i] {
					if tt.args.board[i][j] != tt.res[i][j] {
						t.Fail()
					}
				}
			}
		})
	}
}
