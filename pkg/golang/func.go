package golang

import (
	"fmt"
)

type Book struct {
	pages  int
	cover  int
	author string
}

func (b Book) Pages() int {
	return b.pages
}
func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func call() {
	var book Book

	// 调用这两个隐式声明的函数
	(*Book).SetPages(&book, 123)
	fmt.Println(Book.Pages(book)) // 123
	// 直接调用
	book.SetPages(1)
	fmt.Println(Book.Pages(book)) // 1
}
