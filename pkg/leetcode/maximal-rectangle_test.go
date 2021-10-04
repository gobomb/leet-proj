package leetcode

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				matrix: [][]byte{},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				matrix: [][]byte{
					{'0'},
				},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				matrix: [][]byte{
					{'1'},
				},
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				matrix: [][]byte{
					{'0', '0'},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
