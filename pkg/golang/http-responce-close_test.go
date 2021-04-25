package golang

import (
	"net/http"
	"testing"
	"time"
)

func Test_handler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.w, tt.args.r)
		})
	}
}

func Test_Close(t *testing.T) {
	stopC := make(chan struct{})
	waitC := make(chan struct{})
	timer := time.NewTimer(10 * time.Second)
	go runServer(stopC, waitC)
	<-waitC
	go client()

	<-timer.C

	close(stopC)

}

// func Test_CloseLeak(t *testing.T) {
// 	time.Sleep(2 * time.Second)
// 	stopC := make(chan struct{})
// 	waitC := make(chan struct{})
// 	timer := time.NewTimer(10 * time.Second)
// 	go runServer(stopC, waitC)
// 	<-waitC
// 	// for i := 0; i < 100; i++ {
// 	go clientLeak()
// 	// }

// 	<-timer.C

// 	close(stopC)
// }
