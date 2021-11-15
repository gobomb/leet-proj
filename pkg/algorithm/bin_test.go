package algorithm

import (
	"justest/pkg/utils"
	"testing"
)

type args struct {
	arr []int
	k   int
}

type testBin struct {
	commonTest []struct {
		name string
		args args
		want int
	}
}

var tb testBin

// sorted := genSortedSlices(3)
// log.Printf("%#v", sorted)
// [][]int{[]int{9, 12, 12, 17, 26, 64, 66, 67, 76, 85}, []int{0, 13, 29, 50, 53, 61, 70, 93, 95, 99}, []int{0, 14, 23, 23, 28, 40, 50, 66, 86, 92}}
func init() {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{9, 12, 12, 17, 26, 64, 66, 67, 76, 85},
				76,
			},
			8,
		},
		{
			"2",
			args{
				[]int{0, 13, 29, 50, 53, 61, 70, 93, 95, 99},
				100,
			},
			-1,
		},
		{
			"3",
			args{
				[]int{0, 14, 23, 23, 28, 40, 50, 66, 86, 92},
				15,
			},
			-1,
		},
		{
			"4",
			args{
				[]int{0, 14, 23, 23, 28, 40, 50, 66, 86, 92},
				92,
			},
			9,
		},
		{
			"",
			args{
				[]int{0},
				0,
			},
			0,
		},
	}
	tb = testBin{
		commonTest: tests,
	}
}

func TestBin(t *testing.T) {
	testFunc := []func([]int, int) int{
		BinSearch,
		BinSearch1,
	}

	for _, f := range testFunc {
		for _, tt := range tb.commonTest {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.arr, tt.args.k); got != tt.want {
					t.Errorf("%v() = %v, want %v", utils.FuncName(f), got, tt.want)
				}
			})
		}
	}
}

func TestBinSearch2(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   3,
			},
			want: 5,
		},
	}

	tests = append(tests, tb.commonTest...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinSearch2(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("BinSearch2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinSearch1(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   7,
			},
			want: -1,
		},
	}

	tests = append(tests, tb.commonTest...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinSearch1(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("BinSearch1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinSearch3(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   3,
			},
			want: 6,
		},
		{
			name: "大于全部的数",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   8,
			},
			want: -1,
		},
		{
			name: "小于全部的数",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   -10,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinSearch3(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("BinSearch3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinSearch4(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   3,
			},
			want: 1,
		},
		{
			name: "大于全部的数",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   8,
			},
			want: 8,
		},
		{
			name: "小于全部的数",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   -10,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   2,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinSearch4(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("BinSearch4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumberOfK(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   3,
			},
			want: 4,
		},
		{
			name: "大于全部的数",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   8,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5, 6},
				k:   2,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfK(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("getNumberOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}
