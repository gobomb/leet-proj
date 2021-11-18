package priorityqueue

import (
	"log"
	"testing"
)

func Test_pq(t *testing.T) {
	p := Pq(100)
	// log.Println(p.elem)
	// p.Insert(5)
	// // p.print()
	// // log.Println(p.elem)

	// p.Insert(10)
	// p.Insert(3)
	// log.Println(p.elem)
	// p.DeleteMin()
	// log.Println(p.elem)
	// p.DeleteMin()
	// log.Println(p.elem)
	in := []int{13, 14, 19, 65, 26, 32, 31, 21, 19, 68, 16}
	for _, ii := range in {
		p.Insert(ii)
	}
	p.print()
	p.DeleteMin()
	p.print()
	rs := []int{}
	for !p.IsEmpty() {
		k := p.DeleteMin()
		p.print()
		rs = append(rs, k)
		if k == -1 {
			break
		}
	}

	log.Println(rs)
}
