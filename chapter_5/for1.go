package main

import (
	"fmt"
)

func main() {
	// for基本形式 -> for 初始化语句; 条件语句; 修饰语句 {}
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v) //==>00000
		v = 5
	}

	// 死循环
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
		break
	}
	// 死循环
	for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
		break
	}

	// s := ""
	// for s != "aaaaa" {
	//     fmt.Println("Value of s:", s)
	//     s = s + "a" //
	// }
	// 您还可以在循环中同时使用多个计数器：

	// for i, j := 0, N; i < j; i, j = i+1, j-1 {}
	// 这得益于 Go 语言具有的平行赋值的特性
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" { //for变量在for的block里生效
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
