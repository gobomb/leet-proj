package algorithm

import (
	"justest/pkg/utils"
	"reflect"
	"sort"
	"testing"
)

func Test_mergeTwoSortList(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	tests := func() []struct {
		name string
		args args
		want []int
	} {

		sts := []struct {
			name string
			args args
			want []int
		}{}

		for i := 0; i < 10; i++ {
			st := struct {
				name string
				args args
				want []int
			}{
				name: "",
				args: args{
					a: utils.GenSortedSlices(1)[0],
					b: utils.GenSortedSlices(1)[0],
				},
				want: []int{},
			}
			c := append(st.args.a, st.args.b...)
			sort.Ints(c)
			st.want = c
			sts = append(sts, st)
		}

		return sts
	}

	for _, tt := range tests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoSortList(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := func() []struct {
		name string
		args args
		want []int
	} {

		sts := []struct {
			name string
			args args
			want []int
		}{}

		rand, sort := utils.GenRandSlice(10, 100)
		for i := 0; i < 10; i++ {
			st := struct {
				name string
				args args
				want []int
			}{
				name: "",
				args: args{
					arr: rand,
				},
				want: sort,
			}
			sts = append(sts, st)
		}

		return sts
	}

	for _, tt := range tests() {
		t.Run(tt.name, func(t *testing.T) {
			mergeSort(tt.args.arr)
		})
	}
}
