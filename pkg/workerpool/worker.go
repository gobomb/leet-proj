package workerpool

import (
	"os"

	// "google.golang.org/appengine/log"
	"fmt"
)

var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// 用来写 job，job 的最初来源
// A buffered channel that we can send work requests on.
var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, 100)
}

// Worker represents the worker that executes the job
type Worker struct {
	// 是公共的
	WorkerPool chan chan Job
	// 用来获取 job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// 把可用的 job channel 放到 workerpool 里
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Payload.UploadToS3(); err != nil {
					fmt.Printf("Error uploading to S3: %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		// log.Println("stopping worker")
		w.quit <- true
	}()
}
