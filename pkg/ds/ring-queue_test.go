package ds

import (
	"fmt"
)

func ExampleQueue() {
	var q queue = NewRingQueue(5)

	q.push(1)
	fmt.Println(q.list())

	q.push(2)
	q.push(3)
	fmt.Println(q.list())

	q.push(4)
	q.push(5)
	fmt.Println(q.list())

	q.pop()

	fmt.Println(q.list())

	q.pop()

	fmt.Println(q.list())

	q.push(6)
	q.push(7)

	fmt.Println(q.list())
	q.pop()
	fmt.Println(q.list())

	q.push(8)

	fmt.Println(q.list())

	//OutPut:
	//[1]
	//[1 2 3]
	//[1 2 3 4 5]
	//[2 3 4 5]
	//[3 4 5]
	//[3 4 5 6 7]
	//[4 5 6 7]
	//[4 5 6 7 8]
}
