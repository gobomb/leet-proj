package switchCase

import (
	"log"
)

func switchItf(itf interface{}) {
	switch itf := itf.(type) {
	case []int:
		log.Printf("int: %d", itf[0])
	case []string:
		log.Printf("string: %s", itf[0])
	}
}
