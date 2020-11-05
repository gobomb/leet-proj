package gpool_test

import (
	// "gpool"
	"runtime"
	"testing"
	"time"
)

func Test_Example(t *testing.T) {
	pool := gpool.New(100)
	println(runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			println(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	println(runtime.NumGoroutine())
}
