package bitmap

import (
	"reflect"
	"testing"
)

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}

	tests := func() []struct {
		name string
		args args
		want bool
	} {
		return []struct {
			name string
			args args
			want bool
		}{
			{
				name: "",
				args: args{
					matchsticks: []int{1, 2, 3, 4, 5, 5},
				},
				want: true,
			},
			{
				name: "",
				args: args{
					matchsticks: []int{1, 2, 3, 4, 5},
				},
				want: false,
			},
			{
				name: "",
				args: args{
					matchsticks: []int{3, 3, 3, 3, 4},
				},
				want: false,
			},
			{
				name: "",
				args: args{
					matchsticks: []int{1, 1, 2, 2, 2},
				},
				want: true,
			},
			{
				name: "",
				args: args{
					matchsticks: []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3},
				},
				want: true,
			},
		}
	}
	toTest := []func([]int) bool{makesquare}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.matchsticks); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
