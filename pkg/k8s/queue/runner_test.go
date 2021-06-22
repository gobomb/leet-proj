package queue

import (
	"errors"
	"log"
	"testing"
	"time"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type dumbCtrl struct {
	handle func(i int)
}

func (d *dumbCtrl) Run(stopCh <-chan struct{}) {
	defer func() {
		log.Printf("%+v\n", "stopped dumbCtrl")
	}()
	i := 0
	for i <= 130 {
		select {
		case <-stopCh:
			return
		default:
			d.handle(i)
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (d *dumbCtrl) HasSynced() bool {
	return true
}

func (d *dumbCtrl) LastSyncResourceVersion() string {
	return ""
}

func TestNewRunner(t *testing.T) {
	type args struct {
		queue        workqueue.RateLimitingInterface
		ctrl         cache.Controller
		threadiness  int
		requeueTimes int
		h            func(interface{}) error
		name         string
	}
	q := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	arg := args{
		queue: q,
		ctrl: &dumbCtrl{
			handle: func(i int) {
				q.Add(i)
			},
		},
		threadiness:  3,
		requeueTimes: 5,
		h: func(obj interface{}) error {
			i, ok := obj.(int)
			if !ok {
				return errors.New("reflect failed")
			}
			log.Printf("%+v\n", i)
			if i > 10 && i < 20 {
				i++
				return errors.New("i>10&&i<20")
			}
			return nil
		},
		name: "test runner",
	}
	tt := struct {
		args args
		name string
	}{
		args: arg,
		name: "test runner",
	}
	t.Run(tt.name, func(t *testing.T) {
		runner := NewRunner(tt.args.queue, tt.args.ctrl, tt.args.threadiness, tt.args.requeueTimes, tt.args.h, tt.args.name)
		timer := time.NewTimer(20 * time.Second)
		stopC := make(chan struct{})
		go runner.Run(stopC)

		for ti := range timer.C {
			log.Printf("%+v\n", ti)
			close(stopC)
			break
		}
		time.Sleep(2 * time.Second)
	})
}
