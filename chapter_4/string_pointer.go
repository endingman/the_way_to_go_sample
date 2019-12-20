package main

import "fmt"

func main() {
	s := "good bye"
	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	var p *string = &s
	// 一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；
	// 你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用。
	*p = "ciao"
	// 指针的格式化标识符为 %p）
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	// 通过对 *p 赋另一个值来更改 “对象”，这样 s 也会随之更改。
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
}
