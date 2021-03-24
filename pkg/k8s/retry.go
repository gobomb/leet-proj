package k8s

import (
	"errors"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

func RetryOnErr(steps, sum int) error {
	i = 0
	lastTime = time.Now()

	err := retry.OnError(wait.Backoff{
		Steps:    steps,
		Duration: 10 * time.Millisecond,
		Factor:   5.0,
		Jitter:   0.1,
	}, func(err error) bool {
		if err != nil {
			log.Printf("in retriable: %v\n", err)
			return true
		}
		return false
	}, func() error {
		return do(sum)
	})
	if err != nil {
		log.Printf("retry err: %v\n", err)
		return err
	}
	return nil
}

var i int
var lastTime time.Time

func do(sum int) error {
	i++
	log.Printf("trying %d...\n", i)
	log.Printf("now aftet last do: %+v\n", time.Now().Sub(lastTime).Seconds())
	lastTime = time.Now()
	if i == sum {
		return nil
	}
	return errors.New("failed")
}
