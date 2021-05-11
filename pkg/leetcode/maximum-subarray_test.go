package leetcode

import (
	"log"
	"testing"
	"time"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
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
	now := time.Now()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
	log.Printf("%v\n", time.Now().Sub(now))

	now = time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray1(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
	log.Printf("%v\n", time.Now().Sub(now))

	now = time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
	log.Printf("%v\n", time.Now().Sub(now))

	now = time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray3(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
	log.Printf("%v\n", time.Now().Sub(now))
}
