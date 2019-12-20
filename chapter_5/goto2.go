// compile error goto2.go:8: goto TARGET jumps over declaration of b at goto2.go:8
package main

import "fmt"

// 如果您必须使用 goto，应当只使用正序的标签（标签位于 goto 语句之后），但注意标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败。
func main() {
	a := 1
	goto TARGET // compile error
	b := 9
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
