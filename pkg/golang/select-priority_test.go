package golang

import (
	"fmt"
	"testing"
	"time"
)

func Test_worker(t *testing.T) {
	type args struct {
		ch1    chan int
		ch2    chan int
		stopCh chan struct{}
	}
	type tst struct {
		name string
		args args
	}
	tests := []tst{
		{
			name: "",
			args: args{
				ch1:    make(chan int),
				ch2:    make(chan int),
				stopCh: make(chan struct{}),
			},
		},
	}

	for _, t := range tests {
		tt := t
		go func(test tst) {

			go func() {
				for i := 0; i < 100; i++ {
					tt.args.ch2 <- 2
				}
			}()

			time.Sleep(100 * time.Microsecond)

			go func() {
				fmt.Println("job1 in")

				for i := 0; i < 100; i++ {
					tt.args.ch1 <- 1
				}
			}()

			_ = time.AfterFunc(5*time.Second, func() {
				tt.args.stopCh <- struct{}{}
			})

		}(tt)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			worker(tt.args.ch1, tt.args.ch2, tt.args.stopCh)
		})
	}
}
