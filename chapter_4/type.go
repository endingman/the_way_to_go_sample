package main

import "fmt"

type TZ int

func main() {
	// 当你在使用某个类型时，你可以给它起另一个名字，然后你就可以在你的代码中使用新的名字（用于简化名称或解决名称冲突）。

	// 在 type TZ int 中，TZ 就是 int 类型的新名称（用于表示程序中的时区），然后就可以使用 TZ 来操作 int 类型的数据

	var a, b TZ = 3, 4 //并行 或 同时 赋值。
	c := a + b
	fmt.Printf("c has the value: %d", c) // 输出：c has the value: 7
}
