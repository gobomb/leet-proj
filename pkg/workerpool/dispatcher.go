package workerpool

import "log"

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	maxWorkers int
	quit       chan struct{}
	workers    []*Worker
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers,
		quit: make(chan struct{})}
}

func (d *Dispatcher) Stop() {
	log.Println("stopping dispatcher")

	d.quit <- struct{}{}

	for _, w := range d.workers {
		w.Stop()
	}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		d.workers = append(d.workers, &worker)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for _, w := range d.workers {
		w.Start()
	}

	for {
		select {
		case job := <-JobQueue:
			// 接收请求会启动大量协程，但这里不做耗时操作，只是分发，不会占用太多内存
			// a job request has been received
			go func(job Job) {
				// 获取一个可用的 job channel
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		case <-d.quit:
			log.Println("stop dispatch")

			return
		}
	}
}
