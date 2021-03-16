package gpool

import (
	"runtime"
	"testing"
	"time"
)

func Test_Example(t *testing.T) {
	pool := New(100)
	t.Log(runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			t.Log(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	t.Log(runtime.NumGoroutine())
}
