package utils

import (
	"fmt"
	"reflect"
)

func DumpMethodSet(i interface{}) {

	v := reflect.TypeOf(i)
	// if v == nil {
	// 	log.Printf("%v is nil, return", i)
	// 	return
	// }

	elemTyp := v.Elem()

	n := elemTyp.NumMethod()

	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s's method set:\n", elemTyp)

	for j := 0; j < n; j++ {
		fmt.Println("-", elemTyp.Method(j).Name)
	}

	fmt.Println()
}
