package sort

import (
	"fmt"
	"justest/pkg/utils"
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	s1 := "string"
	s2 := s1
	s3 := s1[3:]
	fmt.Println(&s1, &s2, &s3)

	return
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

	toTest := []func([]int){bubSort, insertSort, insertSort1, mergeSort,
		countSort, radixSort, bucketSort,
		quick, selectSort,
		shellSort, heapSort,
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
