package main

import "fmt"

type TZ int

/** const Ln2= 0.693147180559945309417232121458\
            176568075500134360255254120680009
**/
// 反斜杠 \ 可以在常量表达式中作为多行的连接符使用

func main() {
	// 声明变量的一般形式是使用 var 关键字：var identifier type
	// 当你在使用某个类型时，你可以给它起另一个名字，然后你就可以在你的代码中使用新的名字（用于简化名称或解决名称冲突）。

	// 在 type TZ int 中，TZ 就是 int 类型的新名称（用于表示程序中的时区），然后就可以使用 TZ 来操作 int 类型的数据

	// 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。记住，所有的内存在 Go 中都是经过初始化的

	var a, b TZ = 3, 4 //并行 或 同时 赋值。
	c := a + b
	fmt.Printf("c has the value: %d", c) // 输出：c has the value: 7
}
