package main

import (
    "fmt"
)
/**
在书写表达式类似 var p *type 时，切记在 * 号和指针名称间留有一个空格，因为 var p*type 是语法正确的，但是在更复杂的表达式中，它容易被误认为是一个乘法表达式！

符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。

对于任何一个变量 var， 如下表达式都是正确的：var == *(&var)。

现在，我们应当能理解 pointer.go 的全部内容及其输出：
 */
func main() {
    var i1 = 5
    fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1) //Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
    //这个地址可以存储在一个叫做指针的特殊数据类型中，在本例中这是一个指向 int 的指针，即 i1：此处使用 *int 表示。
    // 如果我们想调用指针 intP，我们可以这样声明它：
    var intP *int
    intP = &i1 // intP 指向 i1,intP 存储了 i1 的内存地址；它指向了 i1 的位置，它引用了变量 i1
    fmt.Printf("The value at memory location %p is %d\n", intP, *intP)// 符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值 ；这被称为反引用
    //    一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，
    //    在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。
}
