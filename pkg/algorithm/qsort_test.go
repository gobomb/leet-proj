package algorithm

import (
	"fmt"
	"justest/pkg/utils"
	"reflect"
	"sort"
	"testing"
)

func Test_quick(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := func() []struct {
		name string
		args args
		want []int
	} {
		cases := []struct {
			name string
			args args
			want []int
		}{}
		for i := 0; i < 3; i++ {
			r, sr := utils.GenRandSlice(3, 50)
			cases = append(cases, struct {
				name string
				args args
				want []int
			}{"", args{r}, sr})
			t.Logf("%v\n", cases[i])
		}

		arr := []int{15, 25, 25, 18, 17, 48, 12, 25, 2, 8, 15, 20, 27, 46, 48}
		sarr := utils.DeepCopyIntSlice(arr)
		sort.Ints(sarr)

		cases = append(cases, struct {
			name string
			args args
			want []int
		}{"", args{arr}, sarr})

		for i := 0; i < 10; i++ {
			r, sr := utils.GenRandSlice(10, 50)
			cases = append(cases, struct {
				name string
				args args
				want []int
			}{"", args{r}, sr})
		}
		return cases
	}

	toTest := []func([]int){quick, bubSort, selectSort, insertSort, mergeSort, shellSort, heapSort}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				argsPrint := fmt.Sprintf("%v(%+v) ", utils.FuncName(f), tt.args)
				if f(tt.args.arr); !reflect.DeepEqual(tt.args.arr, tt.want) {
					t.Errorf("%v = %+v, want %+v", argsPrint, tt.args.arr, tt.want)
				}
			})
		}
	}
}
