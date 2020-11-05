package panicRecover

import (
	"fmt"
	"log"
)

type A struct {
	B map[string]string
	C string
	D string
}

func Ns() {
	a := &A{
		B: make(map[string]string),
		C: "hell",
		D: "pop",
	}

	a.B["1"] = "2"

	b := a.B

	// a = nil

	log.Println(b["1"])
	fmt.Printf("%+v\n", *a)
	fmt.Printf("%v\n", b)


}
