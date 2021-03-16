package golang

import (
	"fmt"
	"reflect"
	"unsafe"
)

type kk struct{}

func (k kk) test() {
	println("aaa")
}

// 有方法的 sturct{} 内存占用也是0
// 转成interface{} 则为16
func size() {
	var k kk
	var b Book
	// struct{} 都是同一个地址
	println(&k)
	println(&struct{}{})

	println(unsafe.Sizeof(k))  // 0
	println(unsafe.Sizeof(&k)) // 8 go指针占 8 个字节

	println(unsafe.Sizeof(b))  // 32 两个int一个string的大小
	println(unsafe.Sizeof(&b)) // 8 go指针占 8 个字节

	println(unsafe.Sizeof(interface{}(k))) // 16 interface的大小
	println(unsafe.Sizeof(interface{}(b))) // 16

	// println(reflect.TypeOf(k).String())
	// println(reflect.TypeOf(interface{}(k)).String())

	fmt.Printf("%v\n", reflect.ValueOf(1))
	fmt.Printf("%v\n", reflect.ValueOf(k))
	fmt.Printf("%v\n", reflect.ValueOf(interface{}(k)))
}
