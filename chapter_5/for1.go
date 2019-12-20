package main

import (
	"fmt"
)

func main() {

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

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" { //for变量在for的block里生效
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
