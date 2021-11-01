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

func (this *MinStack) Push(val int) {
	top := &no{
		prev:  this.top,
		value: val,
	}
	this.top = top
	this.length++
}

func (this *MinStack) Pop() {
	// if this.length == 0 {
	// }
	top := this.top
	this.top = top.prev
	this.length--
}

func (this *MinStack) Top() int {
	if this.length == 0 {
		return 0
	}
	return this.top.value
}

func (this *MinStack) GetMin() int {
	t := this.top
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
