package main

import (
	"fmt"
)

type B struct {
	thing int
}

/**
change() 接受一个指向 B 的指针，并改变它内部的成员；
write() 接受通过拷贝接受 B 的值并只输出 B 的内容。
注意 Go 为我们做了探测工作，我们自己并没有指出是否在指针上调用方法，Go 替我们做了这些事情。
b1 是值而 b2 是指针，方法都支持运行了。
*/

func (b *B) change() { b.thing = 1 } //*b 指针

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())
}

/* 输出：
{1}
{1}
*/
