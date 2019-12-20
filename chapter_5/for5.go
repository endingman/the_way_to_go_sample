package main

import "fmt"

func main() {
	// 关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件。
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}

	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}
