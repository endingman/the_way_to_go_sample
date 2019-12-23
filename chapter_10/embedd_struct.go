package main

import "fmt"

type A struct {
    ax, ay int
}

/**
 * 内嵌结构体
 * 内层结构体被简单的插入或者内嵌进外层结构体。
 * 这个简单的 “继承” 机制提供了一种方式，使得可以从另外一个或一些类型继承部分或全部实现。
 */
type B struct {
    A
    bx, by float32
}

func main() {
    b := B{A{1, 2}, 3.0, 4.0}
    fmt.Println(b.ax, b.ay, b.bx, b.by)
    fmt.Println(b.A)
}

/**
 * 当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？

外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
例子：

type A struct {a int}
type B struct {a, b int}

type C struct {A; B}
var c C
规则 2：使用 c.a 是错误的，到底是 c.A.a 还是 c.B.a 呢？会导致编译器错误：ambiguous DOT reference c.a disambiguate with either c.A.a or c.B.a。

type D struct {B; b float32}
var d D
规则 1：使用 d.b 是没问题的：它是 float32，而不是 B 的 b。如果想要内层的 b 可以通过 d.B.b 得到。
*/
