package sort

import (
	"fmt"
	"justest/pkg/utils"
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
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

		for i := 0; i < 1; i++ {
			r, sr := utils.GenRandSlice(10, 50)
			cases = append(cases, struct {
				name string
				args args
				want []int
			}{"", args{r}, sr})
		}
		return cases
	}

	toTest := []func([]int){
		bubSort,
		mergeSort,
		countSort, radixSort, bucketSort,
		quick, insertSort, insertSort1,
		selectSort, shellSort,
		heapSort, heapSortSTL,
	}

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
