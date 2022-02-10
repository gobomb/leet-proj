package workerpool

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// https://medium.com/smsjunk/handling-1-million-requests-per-minute-with-golang-f70ac505fcaa

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_WorkerPool(t *testing.T) {
	MaxWorker := 20
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	var content = &PayloadCollection{}
	paylength := 1000
	content.Payloads = make([]Payload, paylength)
	for i := 0; i < paylength; i++ {
		content.Payloads[i].N = i
	}

	jsonStr, err := json.Marshal(content)
	if err != nil {
		log.Panic(err)
	}

	// log.Println(jsonStr)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/payload", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Panic(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(payloadHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		log.Panicf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	} else {
		log.Println(rr.Code)
	}

	// time.Sleep(15 * time.Second)
	// log.Println("stop")

	// dispatcher.Stop()
	time.Sleep(1000 * time.Second)
}
