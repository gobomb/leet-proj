package panicRecover

import (
	"fmt"
	"sync"
)

// defer wg.Wait() 会阻塞 panic
func Wt() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	fmt.Println("main 0")

	func() {
		panic("panic")
	}()
	fmt.Println("main 1")

	wg.Wait()
	fmt.Println("main 2")

}
