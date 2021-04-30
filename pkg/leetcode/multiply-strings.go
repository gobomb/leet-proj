package leetcode

import "fmt"

func multiply(num1 string, num2 string) string {
	product := make([]byte, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			t := int(num1[i]-48) * int(num2[j]-48)
			sum := int(product[j]) + t
			fmt.Printf("sum %v  \n", sum)
			fmt.Printf("sum str %s\n", string(byte(sum+48)))

			// var sum int
			if sum < 10 {
				product[j] = byte(sum + 48)
			} else {
				product[j] = byte(48 + sum%10)
				product[j+1] = byte(48 + sum/10)
			}
		}
	}
	for i := 0; i < len(product); i++ {
		if product[i] == 0 && i+1 < len(product) {
			product = product[i+1:]
		}
	}
	fmt.Printf("%#v\n", product)
	return string(product)
}
