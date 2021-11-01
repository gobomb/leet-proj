package leetcode

type MinStack struct {
	length int
	top    *no
}

type no struct {
	prev  *no
	value int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		length: 0,
		top:    nil,
	}
}

func (ms *MinStack) Push(val int) {
	top := &no{
		prev:  ms.top,
		value: val,
	}
	ms.top = top
	ms.length++
}

func (ms *MinStack) Pop() {
	// if this.length == 0 {
	// }
	top := ms.top
	ms.top = top.prev
	ms.length--
}

func (ms *MinStack) Top() int {
	if ms.length == 0 {
		return 0
	}
	return ms.top.value
}

func (ms *MinStack) GetMin() int {
	t := ms.top
	mi := t.value

	for {
		if t == nil {
			break
		}
		t = t.prev
		mi = min(t.value, mi)
	}
	return mi
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
