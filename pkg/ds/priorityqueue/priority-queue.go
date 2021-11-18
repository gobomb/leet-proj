package priorityqueue

import (
	"justest/pkg/ds/tree"
	"log"
)

type pq struct {
	cap  int
	elem []int
}

func Pq(cap int) *pq {
	p := pq{
		cap:  cap,
		elem: make([]int, 0, cap),
	}
	return &p
}

func (p *pq) Insert(x int) {
	p.elem = append(p.elem, x)
	p.up()
}

func (p *pq) up() {
	for i := len(p.elem) - 1; p.Less(i, i/2); i /= 2 {
		p.Swap(i, i/2)
	}
}

func (p *pq) print() {
	log.Printf("%+v", tree.MakeTree(p.elem))
}

// 交换第一和最后一个元素
// 将第一个元素下滤
// 返回最后一个元素
func (p *pq) DeleteMin() int {
	var child int
	mine := p.elem[0]

	// 交换第一和最后一个元素
	p.Swap(0, p.Len()-1)

	i := 0
	// n是要被返回的元素下表，不可以被交换
	n := p.Len() - 1
	for {
		child = i*2 + 1

		// 一旦 child 越界，则退出循环
		if child >= n || child < 0 {
			break
		}

		// 比较左右子树
		if child+1 < n && p.Less(child+1, child) {
			child++
		}

		// 如果儿子大于当前节点，下滤完成
		if !p.Less(child, i) {
			break
		}

		// 进行交换
		p.Swap(i, child)
		i = child
		// p.print()
	}

	p.elem = p.elem[:len(p.elem)-1]

	return mine
}

func (p *pq) Len() int {
	return len(p.elem)
}

func (p *pq) Less(a, b int) bool {
	return p.elem[a] < p.elem[b]
}

func (p *pq) Swap(a, b int) {
	p.elem[a], p.elem[b] = p.elem[b], p.elem[a]
}

func (p *pq) IsFull() bool {
	return len(p.elem) == p.cap
}

func (p *pq) IsEmpty() bool {
	return len(p.elem) == 0
}
