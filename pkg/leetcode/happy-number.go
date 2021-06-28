package leetcode

func isHappy(nn int) bool {
	var rs int
	n := nn
	cache := make(map[int]struct{})
	cache[n] = struct{}{}

	for n != 0 {
		i := n % 10
		n = n / 10
		rs += i * i

		if n == 0 {
			n, rs = rs, n
			// log.Printf("%+v\n", n)
			if n == 1 {
				return true
			}
			if _, ok := cache[n]; ok {
				return false
			}
			cache[n] = struct{}{}
		}
	}
	return false
}

func isHappy2(nn int) bool {
	var rs int
	n := nn
	for n != 0 {
		i := n % 10
		n = n / 10
		rs += i * i

		if n == 0 {
			n, rs = rs, n
			if n == 1 {
				return true
			}
			if n == 4 {
				return false
			}
		}
	}
	return false
}
