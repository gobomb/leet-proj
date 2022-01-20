package event

import (
	"context"
	"log"
	"sync"
	"time"
)

type key interface{}

type SubPub interface {
	Subscribe(ctx context.Context, key interface{}) chan interface{}
	Publish(ctx context.Context, key, val interface{})
}

func NewEventServer() *EventServer {
	return &EventServer{
		subscribers: map[key]subscriberList{},
	}
}

type EventServer struct {
	subscribers map[key]subscriberList
	guard       sync.Mutex
}

type subscriber struct {
	// sub请求的唯一标志
	id int64
	in chan interface{}
}

type subscriberList []*subscriber

func NewSubscriber() *subscriber {
	return &subscriber{
		time.Now().UnixNano(),
		make(chan interface{}),
	}
}

func (e *EventServer) Subscribe(ctx context.Context, key interface{}) chan interface{} {
	s := NewSubscriber()

	e.guard.Lock()
	e.subscribers[key] = append(e.subscribers[key], s)
	e.guard.Unlock()

	id := s.id

	out := make(chan interface{})
	go func() {
		defer log.Printf("Sub %v closed.", s.id)

		for {
			select {
			case got, ok := <-s.in:
				if ok {
					out <- got
				}
			case <-ctx.Done():
				log.Printf("Timeout or canceled: %v.", ctx.Err().Error())
				close(out)

				e.guard.Lock()
				// 删除subscriber实例
				for i, suber := range e.subscribers[key] {
					if suber.id == id {
						e.subscribers[key][i] = nil
						e.subscribers[key] = append(e.subscribers[key][:i], e.subscribers[key][i+1:]...)
					}
				}

				e.guard.Unlock()

				// 当上一段删除subscriber的代码抢到了锁，删除了subscriber，in channel就不可能会被publish写，
				// 所以这里虽然是读端，也可以安全地关闭 channel
				close(s.in)

				return
			}
		}
	}()

	return out
}

func (e *EventServer) Publish(ctx context.Context, key, val interface{}) {
	e.guard.Lock()
	defer e.guard.Unlock()

	list, ok := e.subscribers[key]

	if ok {
		for _, suber := range list {
			select {
			case <-ctx.Done():
				log.Printf("Timeout or canceled: %v.", ctx.Err().Error())
				return
			default:
			}

			fresh := suber

			select {
			case fresh.in <- val:
			case <-ctx.Done():
				log.Printf("Timeout or canceled: %v.", ctx.Err().Error())
				return
			}
		}
	}
}
