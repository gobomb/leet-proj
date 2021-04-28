package ds

import (
	"crypto/rand"
	"math/big"
	"sort"
)

func randSliceInt(max int, length int) []int {
	rs := []int{}
	for i := 0; i < length; i++ {
		rint, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			panic(err)
		}
		rs = append(rs, int(rint.Int64()))
	}
	return rs
}

func genRandSlice(length, max int) ([]int, []int) {
	r := randSliceInt(max, length)
	rcopy := DeepCopyIntSlice(r)
	sort.Ints(rcopy)
	return r, rcopy
}

func DeepCopyIntSlice(arr []int) []int {
	rcopy := make([]int, len(arr))
	copy(rcopy, arr)
	return rcopy
}
