package panicRecover

import (
	"fmt"
)

func Df(){
	i:=0
	defer func(){
		// if err:=recover(); err!=nil {
		// 	fmt.Printf("recover: %s\n",err)
		// }
		fmt.Printf("defer\n")
		for {
			
		}
	}()
	defer panic(i)
	defer fmt.Printf("%d\n",i)
	i++
	fmt.Printf("end, %d\n", i)

	return
}


