package main

import "fmt"

func main() {

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				// continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}
