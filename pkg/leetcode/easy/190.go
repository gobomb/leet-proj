package easy

/*
	190. Reverse Bits

*/

func reverseBits(num uint32) uint32 {
	var res, n uint32
	var one uint32 = 1

	for i := 0; i < 32; i++ {
		if i*2 < 31 {
			n = num >> (31 - i*2)
		} else {
			n = num << (i*2 - 31)
		}

		res |= (n & (one << i))
	}

	return res
}

// https://leetcode.com/problems/reverse-bits/discuss/426492/Golang-100-super-simple
func reverseBits1(num uint32) uint32 {
	var res uint32

	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}

	return res
}
