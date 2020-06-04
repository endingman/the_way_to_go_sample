package main

import (
	"fmt"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

/**
  改变外部变量（outside variable）
  传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回。
  如下的例子，reply 是一个指向 int 变量的指针，通过这个指针，我们在函数内修改了这个 int 变量的数值。
  仅仅是个指导性的例子，当需要在函数内改变一个占用内存比较大的变量时，性能优势就更加明显了
* @Author Liumm
* @Date   2019-08-19
* @return {[type]}   [description]
*/
func main() {
	n := 0
	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	// var i1 = 5
	// fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	reply := &n
	fmt.Println("Multiply1:", *reply)
	Multiply(10, 5, reply)
	fmt.Println("Multiply2:", *reply) // Multiply: 50
}
