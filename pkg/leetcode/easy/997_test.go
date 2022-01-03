package easy

import (
	"reflect"
	"testing"
)

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}

	tests := func() []struct {
		name string
		args args
		want int
	} {
		return []struct {
			name string
			args args
			want int
		}{
			{
				name: "",
				args: args{
					n:     2,
					trust: [][]int{{1, 2}},
				},
				want: 2,
			},
			{
				name: "",
				args: args{
					n:     3,
					trust: [][]int{{1, 3}, {2, 3}},
				},
				want: 3,
			},
			{
				name: "",
				args: args{
					n:     3,
					trust: [][]int{{1, 2}, {2, 3}},
				},
				want: -1,
			},
			{
				name: "",
				args: args{
					n:     3,
					trust: [][]int{{1, 3}, {2, 3}, {3, 1}},
				},
				want: -1,
			},
			{
				name: "",
				args: args{
					n:     1,
					trust: [][]int{},
				},
				want: 1,
			},
		}
	}
	toTest := []func(int, [][]int) int{findJudge, findJudge1}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.n, tt.args.trust); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
