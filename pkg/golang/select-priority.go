package golang

import (
	"fmt"
)

func worker(ch1, ch2 <-chan int, stopCh chan struct{}) {
	// 使job1优先执行
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			// 清空 ch1 后再处理 job2
			fmt.Println(job2)
		}
	}
}
