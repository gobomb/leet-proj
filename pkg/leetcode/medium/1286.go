package medium

/*
	1286. Iterator for Combination
*/

type CombinationIterator struct {
	cache []string
	cur   int
}

/*
	Runtime: 8 ms, faster than 100.00% of Go online submissions for Iterator for Combination.
	Memory Usage: 7.6 MB, less than 12.50% of Go online submissions for Iterator for Combination.
*/
func Constructor(characters string, combinationLength int) CombinationIterator {
	return CombinationIterator{
		cache: makePermutation(characters, 0, combinationLength, []byte{}),
	}
}

func makePermutation(s string, ind int, length int, bs []byte) (res []string) {
	if len(bs) == length {
		res = append(res, string(bs))
		return
	}

	for i := ind; i < len(s); i++ {
		bs = append(bs, s[i])
		res = append(res, makePermutation(s, i+1, length, bs)...)
		bs = bs[:len(bs)-1]
	}

	return
}

func (ci *CombinationIterator) Next() string {
	if !ci.HasNext() {
		return ""
	}

	ci.cur++

	return ci.cache[ci.cur-1]
}

func (ci *CombinationIterator) HasNext() bool {
	return ci.cur < len(ci.cache)
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
