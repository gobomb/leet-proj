package escap_analysis

import "fmt"

func Ea() {
	a := new(int)
	fmt.Println("value of a: ", a)
	b := new(int)
	sum := *a + *b
	fmt.Println("sum = ", sum)
}
