package event

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_newEventServer(t *testing.T) {
	es := NewEventServer()

	wg := sync.WaitGroup{}
	wg.Add(2)

	// 使pub在sub准备好后才运行
	subed := sync.WaitGroup{}
	subed.Add(2)

	fn := func(key string) {
		defer wg.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		ch := es.Subscribe(ctx, key)

		subed.Done()

		for {
			select {
			case got, ok := <-ch:
				if !ok {
					log.Println("Chan closed.")
					return
				}
				log.Printf("Got msg %v on topic %v", got, key)
			}
		}

	}

	go fn("k")
	go fn("j")

	subed.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	es.Publish(ctx, "k", 1)
	es.Publish(ctx, "k", 12)
	es.Publish(ctx, "k", 13)
	es.Publish(ctx, "k", 14)
	es.Publish(ctx, "j", 3)

	wg.Wait()
	// log.Println(es.subscribers)
}

func init() {
	log.SetFlags(log.Ldate | log.Llongfile)
}
