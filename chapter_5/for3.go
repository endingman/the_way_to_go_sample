package main

import (
	"fmt"
)

// 关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件。
// continue只能在for里使用
func main() {
	// 一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）。但在 switch 或 select 语句中（详见第 13 章），break 语句的作用结果是跳过整个代码块，执行后续的代码。

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}

	j := 0
	for { //since there are no checks, this is an infinite loop
		if j >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of j is:", j)
		j++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
