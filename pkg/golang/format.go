package golang

import (
	"fmt"
)

// https://github.com/gobomb/myDoc/issues/9#issue-496928761

type book struct {
	page int
	indp *index
	inds index
}

type index struct {
	count int
}

/*
v  struct   {1 0xc0000181b0 {111}}
v  pointer  &{1 0xc0000181b0 {111}}
+v struct   {page:1 indp:0xc0000181b0 inds:{count:111}}
+v pointer  &{page:1 indp:0xc0000181b0 inds:{count:111}}
#v struct   MY #v golang.index{count:11}
#v pointer  MY #v golang.index{count:11}
*/
// func (b book) GoString() string {
// 	return fmt.Sprintf("MY #v %#v", *b.indp)
// }

/*
v  struct   MY v & +v golang.index{count:11}
v  pointer  MY v & +v golang.index{count:11}
+v struct   MY v & +v golang.index{count:11}
+v pointer  MY v & +v golang.index{count:11}
#v struct   golang.book{page:1, indp:(*golang.index)(0xc0001240b8), inds:golang.index{count:111}}
#v pointer  &golang.book{page:1, indp:(*golang.index)(0xc0001240b8), inds:golang.index{count:111}}
*/
// func (b book) String() string {
// 	return fmt.Sprintf("MY v & +v %#v", *b.indp, )
// }

/*
v  struct   {1 0xc0000181b0 {111}}
v  pointer  MY v & +v golang.index{count:11}
+v struct   {page:1 indp:0xc0000181b0 inds:{count:111}}
+v pointer  MY v & +v golang.index{count:11}
#v struct   golang.book{page:1, indp:(*golang.index)(0xc0000181b0), inds:golang.index{count:111}}
#v pointer  &golang.book{page:1, indp:(*golang.index)(0xc0000181b0), inds:golang.index{count:111}}
*/

// func (b *book) String() string {
// 	return fmt.Sprintf("MY v & +v %#v", *b.indp)
// }

/*
v  struct   {1 0xc0001180b8 {111}}
v  pointer  &{1 0xc0001180b8 {111}}
+v struct   {page:1 indp:0xc0001180b8 inds:{count:111}}
+v pointer  &{page:1 indp:0xc0001180b8 inds:{count:111}}
#v struct   golang.book{page:1, indp:(*golang.index)(0xc0001180b8), inds:golang.index{count:111}}
#v pointer  &golang.book{page:1, indp:(*golang.index)(0xc0001180b8), inds:golang.index{count:111}}
*/
func format() {
	bs := book{1, &index{11}, index{111}}
	switchInterface(bs)
	switchInterface(&bs)

	fmt.Printf("v  struct   %v\n", bs)
	fmt.Printf("v  pointer  %v\n", &bs)
	fmt.Printf("+v struct   %+v\n", bs)
	fmt.Printf("+v pointer  %+v\n", &bs)
	fmt.Printf("#v struct   %#v\n", bs)
	fmt.Printf("#v pointer  %#v\n", &bs)
}

func switchInterface(b interface{}) {
	// type swtich 不能使用 fallthrough
	switch b.(type) {
	case *book:
		fmt.Println("is *book")
	case book:
		fmt.Println("is book")
	default:
		fmt.Println("is not *book nor bokk")
	}
	switch b.(type) {
	case fmt.Stringer:
		fmt.Println("is Stringer")
	default:
		fmt.Println("is not Stringer")
	}
	switch b.(type) {
	case fmt.GoStringer:
		fmt.Println("is GoStringer")
	default:
		fmt.Println("is not GoStringer")
	}
}
