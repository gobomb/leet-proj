package medium

import (
	"fmt"
)

func ExamplePeekingIterator() {
	peekingIterator := peekingConstructor(&Iterator{
		sl: []int{1, 2, 3},
	}) // [1,2,3]

	rs := peekingIterator.peek()
	fmt.Println(rs)

	fmt.Println(peekingIterator.next())
	fmt.Println(peekingIterator.peek())
	fmt.Println(peekingIterator.next())
	fmt.Println(peekingIterator.next())
	// OutPut: 1
	// 1
	// 2
	// 2
	// 3
}
