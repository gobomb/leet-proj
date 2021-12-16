package gpool

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Pool interface {
	Add(int)
	Done()
	Wait()
}

type poolSema struct {
	sema *semaphore.Weighted
	size int64
}

func NewPoolSema(size int) Pool {
	if size <= 0 {
		size = 1
	}
	return &poolSema{
		semaphore.NewWeighted(int64(size)),
		int64(size),
	}
}

func (p *poolSema) Add(delta int) {
	p.sema.Acquire(context.Background(), int64(delta))
}
func (p *poolSema) Done() {
	p.sema.Release(1)
}
func (p *poolSema) Wait() {
	p.sema.Acquire(context.Background(), p.size)
}

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) Pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}
