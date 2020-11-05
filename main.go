package main

// type Book struct {
// 	pages int
// }

// func (b Book) Pages() int {
// 	return b.pages
// }
// func (b *Book) SetPages(pages int) {
// 	b.pages = pages
// }

// type kk struct{}

// func (k kk) test() {
// 	println("aaa")
// }

// func main() {
// 	var book Book

// 	var k kk
// 	t := unsafe.Sizeof(k)
// 	println(t)
// 	// 调用这两个隐式声明的函数。
// 	(*Book).SetPages(&book, 123)
// 	fmt.Println(Book.Pages(book)) // 123

// 	book.SetPages(1)
// }

import (
	"fmt"
	"justest/pkg/ds"
)

//go:noinline
func foo(m0 int) *int {
	var m1 int = 11
	var m2 int = 12
	var m3 int = 13
	var m4 int = 14
	var m5 int = 15

	println(&m0, &m1, &m2, &m3, &m4, &m5)

	return &m3
}

func main() {

	fmt.Println(10>>2)
	ds.TestBin()
	m := foo(100)
	println(*m)
}
