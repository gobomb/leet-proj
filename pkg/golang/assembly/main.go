package main


// go tool compile -S -N -l main.go > main.o
func main() {
	add(3, 5)
}

func add(a, b int) int {
	var t int
	t = 2
	_ = t
	return a + b
}
