package easy

import (
	"log"
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()
	log.Println(param_2, param_3, param_4)
}
