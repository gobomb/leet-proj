package main

import (
	"log"
	"sync"
)

// 《Go语言高级编程》6.2
type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}

	return l
}

func (l Lock) Lock() bool {
	lockResult := false

	select {
	case <-l.c:
		lockResult = true
	default:
		log.Println("default", lockResult)
	}
	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int

func main() {
	var l = NewLock()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for !l.Lock() {

				println("lock failed")
				// return
			}
			// time.Sleep(1 * time.Second)
			counter++

			println("currnt counter ", counter)

			l.Unlock()
		}()
	}
	wg.Wait()

}
