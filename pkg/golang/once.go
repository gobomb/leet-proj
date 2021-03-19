package golang

import "sync"

// https://mp.weixin.qq.com/s/ELnUQL-EnK8w8EpdbZKDWA
func once() {
	o := sync.Once{}
	o.Do(func() { println("do") })
}
