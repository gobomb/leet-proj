package queue

import (
	"time"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"k8s.io/klog"

	rt "k8s.io/apimachinery/pkg/util/runtime"
)

type Runner struct {
	queue        workqueue.RateLimitingInterface
	ctrl         cache.Controller
	threadiness  int
	requeueTimes int
	Handler      func(interface{}) error
	name         string
}

type QueueRunner interface {
	Run(stopCh <-chan struct{})
}

func NewRunner(queue workqueue.RateLimitingInterface,
	ctrl cache.Controller, threadiness int,
	requeueTimes int,
	h func(interface{}) error,
	name string) *Runner {
	return &Runner{
		queue,
		ctrl,
		threadiness,
		requeueTimes,
		h,
		name,
	}
}

func (c *Runner) String() string {
	return c.name
}

func (c *Runner) Run(stopCh <-chan struct{}) {
	defer rt.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()
	klog.Infof("Starting controller %v", c)

	go c.ctrl.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, c.ctrl.HasSynced) {
		klog.Errorf("%+v", errors.New("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < c.threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	klog.Infof("Stopped controller %v", c)
}

func (c *Runner) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Runner) processNextItem() bool {
	// Wait until there is a new item in the working queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
	// This allows safe parallel processing because two items with the same key are never processed in
	// parallel.
	defer c.queue.Done(key)

	// Invoke the method containing the business logic
	err := c.Handler(key)
	// Handle the error if something went wrong during the execution of the business logic
	c.handleErr(err, key)
	return true
}

// handleErr checks if an error happened and makes sure we will retry later.
func (c *Runner) handleErr(err error, key interface{}) {
	if err == nil {
		// Forget about the #AddRateLimited history of the key on every successful synchronization.
		// This ensures that future processing of updates for this key is not delayed because of
		// an outdated error history.
		c.queue.Forget(key)
		return
	}

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.queue.NumRequeues(key) < c.requeueTimes {
		klog.Infof("Error syncing item %+v: %+v", key, err)

		// Re-enqueue the key rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the key will be processed later again.
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	// Report to an external entity that, even after several retries, we could not successfully process this key
	// rt.HandleError(err)

	klog.Infof("Dropping item %q out of the queue: %+v", key, err)
}
