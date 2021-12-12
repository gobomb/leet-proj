package medium

import (
	"log"
	"testing"
)

func Test_peekingConstructor(t *testing.T) {
	peekingIterator := peekingConstructor(&Iterator{
		sl: []int{1, 2, 3},
	}) // [1,2,3]
	var rs int
	rs = peekingIterator.peek() // return 2, the pointer does not move [1,2,3].
	log.Println(rs)

	rs = peekingIterator.next() // return 1, the pointer moves to the next element [1,2,3].
	log.Println(rs)

	rs = peekingIterator.peek() // return 2, the pointer does not move [1,2,3].
	log.Println(rs)

	rs = peekingIterator.next() // return 2, the pointer moves to the next element [1,2,3]
	log.Println(rs)

	rs = peekingIterator.next() // return 3, the pointer moves to the next element [1,2,3]
	log.Println(rs)

	// return False
	log.Println(peekingIterator.hasNext())

}
