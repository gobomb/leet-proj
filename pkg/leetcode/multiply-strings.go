package leetcode

func tobyte(i int) byte {
	return byte(48 + i)
}

func toint(b byte) int {
	if b == 0 {
		return 0
	}
	return int(b - '0')
}

func multiply(num1 string, num2 string) string {
	product := make([]byte, len(num1)+len(num2))

	for i := 0; i < len(num1); i++ {
		ni := len(num1) - 1 - i
		for j := 0; j < len(num2); j++ {
			nj := len(num2) - 1 - j
			nk := len(product) - 1 - (i + j)
			prd := toint(num1[ni]) * toint(num2[nj])
			sum := toint(product[nk]) + prd

			for sum >= 10 {
				product[nk] = tobyte(sum % 10)
				sum /= 10
				nk--
				sum += toint(product[nk])
			}

			if sum < 10 {
				product[nk] = tobyte(sum % 10)
			}
		}
	}

	k := 0

	for k = 0; k < len(product); k++ {
		if toint(product[k]) != 0 {
			break
		}
	}

	if k == len(product) {
		product = product[k-1:]
	} else {
		product = product[k:]
	}

	return string(product)
}
