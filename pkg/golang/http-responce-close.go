package golang

import (
	// "fmt"

	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// w.Header().Set("Location", "http://0.0.0.0:8080")
	// w.WriteHeader(http.StatusMovedPermanently)
	// w.Write([]byte("hello"))

	// 递归重定向到自己，使客户端发生错误
	http.Redirect(w, r, "http://0.0.0.0:8080", http.StatusMovedPermanently)
}

func runServer(stopC <-chan struct{}, waitC chan<- struct{}) {
	var hf http.HandlerFunc = handler
	var myServeMux = http.NewServeMux()

	myServeMux.HandleFunc("/", hf)
	server := &http.Server{Addr: ":8080", Handler: myServeMux}

	go func() {
		<-stopC
		server.Close()
		log.Println("stopC resievd")

	}()
	log.Printf("run http server at :8080\n")
	waitC <- struct{}{}
	log.Println(server.ListenAndServe())
	log.Println("server closed")
}

func client() {
	client := &http.Client{}
	url := "http://0.0.0.0:8080/"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

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
	// 10次重定向会报错：stopped after 10 redirects
	if err != nil {
		log.Printf("get err:%v\n", err)
		return
	}

	log.Printf("client end!\n")

}

func clientLeak() {
	client := &http.Client{}
	// 可以自定义重定向的行为，默认是自动重定向至多10次
	// client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
	// 	return http.ErrUseLastResponse
	// }
	url := "http://0.0.0.0:8080/"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(reqest)
	if err != nil {
		log.Printf("get err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("client end!\n")
}
