package easy

/*
	232. Implement Queue using Stacks
*/

type MyQueue struct {
	s1 []int // to push
	s2 []int // to pop
}

func ConstructorQueue() (mq MyQueue) {
	return
}

func (mq *MyQueue) Push(x int) {
	mq.s1 = append(mq.s1, x)
}

func (mq *MyQueue) Pop() int {
	if mq.Empty() {
		return -1
	}

	if len(mq.s2) != 0 {
		rs := mq.s2[len(mq.s2)-1]
		mq.s2 = mq.s2[:len(mq.s2)-1]
		return rs
	}

	for i := len(mq.s1) - 1; i >= 0; i-- {
		mq.s2 = append(mq.s2, mq.s1[i])
	}
	mq.s1 = []int{}

	rs := mq.s2[len(mq.s2)-1]
	mq.s2 = mq.s2[:len(mq.s2)-1]
	return rs
}

func (mq *MyQueue) Peek() int {
	if mq.Empty() {
		return -1
	}

	if len(mq.s2) != 0 {
		return mq.s2[len(mq.s2)-1]
	}

	for i := len(mq.s1) - 1; i >= 0; i-- {
		mq.s2 = append(mq.s2, mq.s1[i])
	}
	mq.s1 = []int{}

	return mq.s2[len(mq.s2)-1]
}

func (mq *MyQueue) Empty() bool {
	return len(mq.s1) == 0 && len(mq.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// 6 7 8 9
//
