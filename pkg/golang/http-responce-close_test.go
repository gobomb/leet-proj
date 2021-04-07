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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.w, tt.args.r)
		})
	}
}

func Test_Close(t *testing.T) {
	stopC := make(chan struct{})
	waitC := make(chan struct{})
	timer := time.NewTimer(30 * time.Second)
	go runServer(stopC, waitC)
	<-waitC
	go client()

	<-timer.C
	close(stopC)
}
