package easy

/*
	225. Implement Stack using Queues
*/

type MyStack struct {
	l1 []int
	l2 []int
}

func Constructor() (ms MyStack) {
	return
}

func (ms *MyStack) Push(x int) {
	ms.l1 = append(ms.l1, x)
}

func (ms *MyStack) Pop() int {
	for len(ms.l1) != 1 {
		ms.l2 = append(ms.l2, ms.l1[0])
		ms.l1 = ms.l1[1:]
	}

	rs := ms.l1[0]
	ms.l1 = ms.l1[1:]

	ms.l1, ms.l2 = ms.l2, ms.l1
	return rs
}

func (ms *MyStack) Top() int {
	for len(ms.l1) != 1 {
		ms.l2 = append(ms.l2, ms.l1[0])
		ms.l1 = ms.l1[1:]
	}

	rs := ms.l1[0]

	ms.l2 = append(ms.l2, ms.l1[0])
	ms.l1 = ms.l1[1:]

	ms.l1, ms.l2 = ms.l2, ms.l1
	return rs
}

func (ms *MyStack) Empty() bool {
	return len(ms.l1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
