package medium

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

func (lc *LRUCache) Get(key int) int {
	if v, ok := lc.cache[key]; ok {
		// 从任意位置删除
		lc.Remove(v)
		lc.Add(v)
		return v.val
	}

	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if v, ok := lc.cache[key]; ok {
		v.val = value
		lc.Remove(v)
		lc.Add(v)
	} else {
		node := &DulNode{
			key: key,
			val: value,
		}
		lc.cache[key] = node
		lc.Add(node)
	}

	if len(lc.cache) > lc.Cap {
		delete(lc.cache, lc.tail.key)
		lc.Remove(lc.tail)
	}
}

// 添加到头部
// node的prior和next不一定为nil
func (dl *DulLink) Add(node *DulNode) {
	// 把node的prior清除
	node.prior = nil
	//
	node.next = dl.head
	if dl.head != nil {
		dl.head.prior = node
	}
	dl.head = node
	if dl.tail == nil {
		dl.tail = node
		dl.tail.next = nil
	}
}

// 从任意位置删除
func (dl *DulLink) Remove(node *DulNode) {
	// if node == this.head {
	// 	this.head = nil
	// 	return
	// }
	if node == dl.head {
		dl.head = node.next

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

	if node == dl.tail {
		node.prior.next = nil
		dl.tail = node.prior
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
