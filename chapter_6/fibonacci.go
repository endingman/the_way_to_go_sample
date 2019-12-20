package main

import "fmt"

/**
 * 斐波那数列（递归）
 * @Author Liumm
 * @Date   2019-08-19
 * @return {[type]}   [description]
 */
func main() {
	result := 0
	key := 0
	for i := 0; i <= 10; i++ {
		key, result = fibonacci(i)
		fmt.Printf("fibonacci(%d) and key(%d) is: %d\n", i, key, result)
	}
}

func fibonacci(n int) (key int, res int) {
	if n <= 1 {
		key, res = n+1, 1
	} else {
		key = n + 1
		res1 := 0
		res2 := 0
		_, res1 = fibonacci(n - 1)
		_, res2 = fibonacci(n - 2)
		res = res1 + res2
	}
	return
}
