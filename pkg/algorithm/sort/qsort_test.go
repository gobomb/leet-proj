package sort

import (
	"justest/pkg/utils"
	"log"
	"math/rand"
	"testing"
	"time"
)

func Test_randomPartition(t *testing.T) {
	type args struct {
		arr        []int
		startIndex int
		endIndex   int
	}

	r, _ := utils.GenRandSlice(10, 50)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				arr:        r,
				startIndex: 0,
				endIndex:   len(r) - 1,
			},
		},
		{
			name: "",
			args: args{
				arr:        []int{21, 15, 21, 3, 0, 33, 23},
				startIndex: 0,
				endIndex:   6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())

			got := randomPartition(tt.args.arr, tt.args.startIndex, tt.args.endIndex)

			log.Println(tt.args.arr, got)
		})
	}
}
