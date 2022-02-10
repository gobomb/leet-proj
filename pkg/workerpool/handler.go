package workerpool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	MaxLength int64 = 1000
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// log.Println(r.Body)

	// Read the body into a string for json decoding
	var content = &PayloadCollection{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, content)

	if err != nil {
		log.Println(err)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {
		// log.Printf("read payload %v\n", payload)
		// let's create a job with the payload
		work := Job{Payload: payload}

		// 把 job 送到全局的 queue
		// Push the work onto the queue.
		JobQueue <- work
	}

	log.Println("sent to JobQueue!")

	w.WriteHeader(http.StatusOK)
}
