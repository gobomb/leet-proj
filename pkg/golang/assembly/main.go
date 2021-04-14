package main

// // go tool compile -S -N -l main.go > main.o

// // main1.o
// func main() {
// 	add(3, 5)
// }

// func add(a, b int) int {
// 	var t int
// 	t = 2
// 	_ = t
// 	return a + b
// }

// // main2

// func main() {
// 	add(33)
// }


/*
32 32-40 返回值
24 24-32 入参 a
16 16-24 返回mian的地址
8   8-16 分配 b 时存的 bp
0	0-8 存 b
*/
// func add(a int) int {
// 	b := 44
// 	return a + b
// }
