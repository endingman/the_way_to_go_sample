package main

//布尔型的值只可以是常量 true 或者 false。
func main() {
	// int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
	// uintptr 的长度被设定为足够存放一个指针即可。
	// Go 语言中没有 float 类型。（Go 语言中只有 float32 和 float64）没有 double 类型。
	var a int
	var b int32
	a = 15
	b = a + a // 编译错误，Go 中不允许不同类型之间的混合使用
	b = b + 5 // 因为 5 是常量，所以可以通过编译，允许常量之间的混合使用
}
