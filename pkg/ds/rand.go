package ds

import (
	"crypto/rand"
	"math/big"
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
