package medium

import (
	"fmt"
)

func ExampleConstructorSolution() {
	c := ConstructorSolution([]int{1, 2, 3, 3, 3})
	got := c.Pick(3)
	switch got {
	case 2, 3, 4:
		fmt.Println("true")
	}

	got = c.Pick(1)
	switch got {
	case 0:
		fmt.Println("true")
	}

	got = c.Pick(3)
	switch got {
	case 2, 3, 4:
		fmt.Println("true")
	}

	// OutPut:
	// true
	// true
	// true

}
