package gpool

import (
	"os"
	"os/signal"
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
			t.Log(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	t.Log(runtime.NumGoroutine())
}

func Test_gate(t *testing.T) {
	a, b := 200, 1000
	c := time.Second * 3
	g := gateSth{
		make(gate, a),
		c,
	}

	t.Log(runtime.NumGoroutine())
	for i := 0; i < b; i++ {
		go g.doSth()
	}
	// t.Log(c * time.Duration(b/a+1))
	// time.Sleep(c * time.Duration(b/a+1))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		for runtime.NumGoroutine() >= 5 {
			time.Sleep(c * 2)
			t.Log(runtime.NumGoroutine())
		}

		t.Log("break", runtime.NumGoroutine())

	}()

	<-ch

	t.Log("got notify", runtime.NumGoroutine())
}
