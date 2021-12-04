package easy

import (
	"log"
	"testing"
)

func TestConstructorQueue(t *testing.T) {
	obj := ConstructorQueue()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()

	log.Println(param_2, param_3, param_4)
}
