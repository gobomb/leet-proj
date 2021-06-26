package leetcode

func defangIPaddr(address string) string {
	var addr []byte
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			addr = append(addr, '[', '.', ']')
		} else {
			addr = append(addr, address[i])
		}
	}
	return string(addr)
}
