package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

/**
 结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。

结构体可以包含内嵌结构体。
可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为

Go 语言中的继承是通过内嵌或组合来实现的

*/
type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	fmt.Println("outer is:", outer)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

/**
当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？

外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
例子：

type A struct {a int}
type B struct {a, b int}

type C struct {A; B}
var c C
规则 2：使用 c.a 是错误的，到底是 c.A.a 还是 c.B.a 呢？会导致编译器错误：ambiguous DOT reference c.a disambiguate with either c.A.a or c.B.a。

要使用的话应该是下面方式
c.A.a
c.B.a
*/
