package ds

import (
	"reflect"
	"sort"
	"testing"
)

func Test_quick(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}

	for i := 0; i < 3; i++ {
		r, sr := genRandSlice(3, 50)
		tests = append(tests, struct {
			name string
			args args
			want []int
		}{"", args{r}, sr})
		t.Logf("%v\n", tests[i])
	}
	arr := []int{15, 25, 25, 18, 17, 48, 12, 25, 2, 8, 15, 20, 27, 46, 48}
	sarr := DeepCopyIntSlice(arr)
	sort.Ints(sarr)
	tests = append(tests, struct {
		name string
		args args
		want []int
	}{"", args{arr}, sarr})

	for i := 0; i < 15; i++ {
		r, sr := genRandSlice(15, 50)
		tests = append(tests, struct {
			name string
			args args
			want []int
		}{"", args{r}, sr})
		// t.Logf("%v\n", tests[i])
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quick(DeepCopyIntSlice(tt.args.arr)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quick(%v) = %v, want %v", tt.args.arr, got, tt.want)
			}
		})
	}
}
