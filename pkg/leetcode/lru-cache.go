package leetcode

/*
	146. LRU Cache
*/

type DulNode struct {
	next, prior *DulNode
	key         int
	val         int
}

type LRUCache struct {
	cache map[int]*DulNode
	Cap   int
	DulLink
}

type DulLink struct {
	head, tail *DulNode
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{
		cache:   make(map[int]*DulNode),
		Cap:     capacity,
		DulLink: DulLink{},
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		// 从任意位置删除
		this.Remove(v)
		this.Add(v)
		return v.val
	}

	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.val = value
		this.Remove(v)
		this.Add(v)
	} else {
		node := &DulNode{
			key: key,
			val: value,
		}
		this.cache[key] = node
		this.Add(node)
	}

	if len(this.cache) > this.Cap {
		delete(this.cache, this.tail.key)
		this.Remove(this.tail)
	}
}

// 添加到头部
// node的prior和next不一定为nil
func (this *DulLink) Add(node *DulNode) {
	// 把node的prior清除
	node.prior = nil
	//
	node.next = this.head
	if this.head != nil {
		this.head.prior = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		//
		this.tail.next = nil
		//
	}
}

// 从任意位置删除
func (this *DulLink) Remove(node *DulNode) {
	// if node == this.head {
	// 	this.head = nil
	// 	return
	// }

	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prior = nil
		}
		//把node的next清除
		node.next = nil
		return
	}

	// if node == this.tail {
	// 	node.prior.next = nil
	// 	this.tail = node.prior
	// 	return
	// }

	if node == this.tail {
		node.prior.next = nil
		this.tail = node.prior
		//把node的prior清除
		node.prior = nil
		//
		return
	}

	node.prior.next = node.next
	node.next.prior = node.prior
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
