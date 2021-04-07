package golang

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// w.Header().Set("Location", "http://baidu.com")
	http.Redirect(w, r, "http://0.0.0.0:8080", http.StatusMovedPermanently)
}

func runServer(stopC <-chan struct{}, waitC chan<- struct{}) {
	var hf http.HandlerFunc = handler
	http.HandleFunc("/", hf)
	server := &http.Server{Addr: ":8080", Handler: hf}
	go func() {
		<-stopC
		server.Close()
	}()
	log.Printf("run http server at :8080\n")
	waitC <- struct{}{}
	log.Fatal(server.ListenAndServe())
}

func client() {
	client := &http.Client{}
	url := "http://0.0.0.0:8080/"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	// client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
	// 	return http.ErrUseLastResponse
	// }
	resp, err := client.Do(reqest)
	if resp != nil && err != nil {
		log.Printf("responce and err are not nil...\n")
	}

	if resp != nil {
		defer resp.Body.Close()

		log.Printf("head: %+v\n", resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("read err:%v\n", err)
			// return
			goto CHECKERR
		}
		log.Printf("responce: %s\n", body)
	}
CHECKERR:
	if err != nil {
		log.Printf("get err:%v\n", err)
		return
	}

	log.Printf("client end!\n")

}
