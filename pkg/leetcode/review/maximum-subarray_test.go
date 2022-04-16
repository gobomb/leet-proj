package review

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
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
				"1",
				args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
				6,
			},
			{
				"2",
				args{[]int{1}},
				1,
			},
			{
				"3",
				args{[]int{5, 4, -1, 7, 8}},
				23,
			},
			{
				"4",
				args{[]int{-1}},
				-1,
			},
			{
				"4",
				args{[]int{-1, -2, -3}},
				-1,
			},
		}
	}

	toTest := []func([]int) int{maxSubArray, maxSubArray2, maxSubArray3, maxSubArray4}

	for _, f := range toTest {
		now := time.Now()

		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}

		log.Printf("%v %v\n", funcName(f), time.Since(now))
	}
}
