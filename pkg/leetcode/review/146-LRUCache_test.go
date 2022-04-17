package review

import (
	"log"
	"testing"
)

func TestLruCache(t *testing.T) {
	var null = -1
	var lRUCache LRUCache

	type testData struct {
		opt []string
		num [][]int
		rss []int
	}

	opt := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	num := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	rss := []int{null, null, null, 1, null, -1, null, -1, 3, 4}

	td := []testData{
		{
			opt,
			num,
			rss,
		},
	}
	for _, v := range td {
		opt, num, rss := v.opt, v.num, v.rss
		for i := range opt {
			switch opt[i] {
			case "LRUCache":
				lRUCache = ConstructorLRU(num[i][0])
			case "put":
				lRUCache.Put(num[i][0], num[i][1])
			case "get":
				rs := lRUCache.Get(num[i][0])
				if rs != rss[i] {
					t.Errorf("failed at want %v but got %v", num[i][0], rs)
				}
			}

			log.Printf("%+v,%v,%v", lRUCache.cache, lRUCache.head, lRUCache.tail)
		}
	}
}
