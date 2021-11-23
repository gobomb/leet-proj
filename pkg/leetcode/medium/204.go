package medium

import (
	"log"
)

/*
	204. Count Primes
*/

// Last executed input 499979
func countPrimes(n int) int {
	counter := func(c chan int) {
		i := 2
		for {
			c <- i
			i++
		}
	}

	filter := func(prime int, recv chan int, send chan int) {
		for i := range recv {
			if i%prime != 0 {
				send <- i
			}
		}
	}

	output := make(chan int)
	c := make(chan int)

	go counter(c)

	go func() {
		for i := 0; i <= n; i++ {
			p := <-c
			log.Println(p)

			output <- p
			primes := make(chan int)
			go filter(p, c, primes)
			c = primes
		}
	}()

	res := 0

	for i := range output {
		if i >= n {
			return res
		}

		res++
	}
	return res
}

// https://leetcode.com/problems/count-primes/discuss/837745/Golang-4ms-beats-97.5
func countPrimes2(n int) int {
	if n <= 2 {
		return 0
	}

	pp := make([]bool, n)
	// 计算了2
	c := 1

	for i := 0; i < n; i++ {
		pp[i] = true
	}

	// 2的倍数必然不是素数，所以i+=2
	for i := 3; i < n; i += 2 {
		if pp[i] {
			c++
			// 只验证j的非2倍乘积
			for j := i * i; j < n; j += 2 * i {
				pp[j] = false
			}
		}
	}

	return c
}
