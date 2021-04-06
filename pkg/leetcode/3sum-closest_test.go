package leetcode

import (
	"testing"
)

type args struct {
	nums   []int
	target int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"1", args{nums: []int{-1, 2, 1, -4}, target: 1}, 2},
	{"2", args{nums: []int{-1, 2, 1, 4}, target: 1}, 2},
	{"3", args{nums: []int{2, 2, 2, 2}, target: 1}, 6},
	{"4", args{nums: []int{1, 1, 1, 0}, target: 100}, 3},
	// TODO: Add test cases.
}

func Test_threeSumClosest(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSumClosestDoublePointer(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosestDoublePointer(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosestDoublePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Benchmark_threeSumClosestDoublePointer(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		b.ReportAllocs()
// 		for _, tt := range tests {
// 			if got := threeSumClosestDoublePointer(tt.args.nums, tt.args.target); got != tt.want {
// 			}
// 			break
// 		}
// 	}
// }

// func Benchmark_threeSumClosest(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		b.ReportAllocs()
// 		for _, tt := range tests {
// 			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
// 			}
// 			break
// 		}
// 	}
// }
