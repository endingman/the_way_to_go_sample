package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	// 变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的。

	// 这可用于在返回语句之后修改返回的 error 时使用。
	fmt.Println(f())
}
