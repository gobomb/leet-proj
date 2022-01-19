package observer

func ExampleObserver() {
	shirtItem := NewItem("Nike Shirt")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}
	observer3 := &customer{id: "cde@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.Register(observer3)
	shirtItem.NotifyAll()
	shirtItem.Deregister(observer3.id)
	shirtItem.NotifyAll()

	// OutPut:
}
