package gpool

import (
	"log"
	"runtime"
	"time"
)

// 	"golang.org/x/tools/godoc/vfs/gatefs"

/*
	不仅可以控制最大的并发数目，
	而且可以通过带缓存通道的使用量和最大容量比例来判断程序运行的并发率。
	当通道为空时可以认为是空闲状态，当通道满了时可以认为是繁忙状态
*/

func init() {
	log.SetFlags(log.Lshortfile)
}

type gate chan bool

func (g gate) enter() { g <- true }
func (g gate) leave() { <-g }

type gateSth struct {
	gate
	sleep time.Duration
}

// 同时进入doSth的只有 len(g.gate)，其他goroutine会被阻塞
func (g gateSth) doSth() {
	g.enter()
	defer g.leave()

	log.Println(runtime.NumGoroutine(), len(g.gate))
	time.Sleep(g.sleep * time.Second)
	// ...
	// fmt.Println("do sth")
}
