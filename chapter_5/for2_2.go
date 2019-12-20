package main

import "fmt"

func main() {
	var i int = 5
	// for 结构的第二种形式是没有头部的条件判断迭代（类似其它语言中的 while 循环），基本形式为：for 条件语句 {}

	//条件语句是可以被省略的，如 i:=0; ; i++ 或 for { } 或 for ;; { }（;; 会在使用 gofmt 时被移除）：这些循环的本质就是无限循环。最后一个形式也可以被改写为 for true { }，但一般情况下都会直接写 for { }。
	//
	//想要直接退出循环体，可以使用 break 语句（第 5.5 节）或 return 语句直接返回（第 6.1 节）。

	// 但这两者之间有所区别，break 只是退出当前的循环体，而 return 语句提前对函数进行返回，不会执行后续的代码。

	// 无限循环的经典应用是服务器，用于不断等待和接受新的请求
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}
}
