package panicRecover

import (
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"
)

// runtime.HandleCrash 会打印两次 panic 的 callstack
func Pr() {
	// defer func() {
	// 	r := recover()
	// 	if r != nil {
	// 		log.Println("recover in main!")
	// 		time.Sleep(10 * time.Minute)
	// 	}
	// }()
	var nilstrust *fakeStruct

	go func() {
		defer runtime.HandleCrash(func(r interface{}) {
			log.Printf("into recover panic!\n")
			// runtime.ReallyCrash = false
		})
		defer func() {
			time.Sleep(10 * time.Second)
		}()
		defer func() {
			r := recover()
			if r != nil {
				log.Println("recover in gofunc!")
				// time.Sleep(10 * time.Minute)
				panic("panic again")
			}
		}()
		log.Printf("In defer")
		fakeFun(nilstrust.s[1])
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Println("hhhhhhhhhhh")
		}
	}()

	log.Printf("Out of defer")

	select{}
}

type fakeStruct struct {
	s []string
}

func fakeFun(s string) {

}
