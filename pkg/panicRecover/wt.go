package panicRecover

import (
	"fmt"
	"sync"
)

func Wt() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// go func() {
	// 	defer wg.Wait()
	// 	// Long running task
	// 	panic("Something unexpected happened.")
	// 	fmt.Println("wait")
	// }()
	// go func() {
	// 	wg.Done()
	// 	fmt.Println("done")

	// }()
	defer wg.Wait()
	func() {
		panic("panic")
		fmt.Println("done")

	}()
	fmt.Println("main 1")

	wg.Wait()
	fmt.Println("main 2")

}
