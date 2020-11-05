package switchCase

import (
	"log"
)

func sc() {
	it := []int{1, 2, 3}
	st := []string{"a", "b"}
	SwitchItf(interface{}(it))
	SwitchItf(interface{}(st))
}

func switchItf(itf interface{}) {
	switch itf := itf.(type) {
	case []int:
		log.Printf("int: %d", itf[0])
	case []string:
		log.Printf("string: %s", itf[0])
	}
}
