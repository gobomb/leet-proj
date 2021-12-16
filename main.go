package main

import (
	"fmt"
	"sync"
	"time"
)

type IA interface {
	fa()
	fb(int) int
}

type A struct{}

func (A) fa() {
	fmt.Println("call fa")
}
func (A) fb(a int) int {
	return 1
}

type B struct {
	IA
}

func (B) fa() {
	fmt.Println("call b fa")
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

func main1() {

	var ia IA
	ia = B{
		A{},
	}

	ia.fa()

	chan1 := make(chan int)
	chan2 := make(chan int)

	writeFlag := false
	go func() {
		for {
			if writeFlag {
				chan1 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			if writeFlag {
				chan2 <- 1
			}
			time.Sleep(time.Second)
			close(chan2)
		}
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")
}
