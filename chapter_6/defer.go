package main

import "fmt"

func main() {
	function1()
	// a()
	f()
}

/**
 * 关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。

 * @Author Liumm
 * @Date   2019-08-19
 * @return {[type]}   [description]
*/
func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!\n")
}

// func a() {
//     i := 0
//     defer fmt.Println(i)
//     i++
//     return
// }

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

/**
 * 1、关闭文件流 （详见 第 12.2 节）
// open a file
defer file.Close()
2、解锁一个加锁的资源 （详见 第 9.3 节）
mu.Lock()
defer mu.Unlock()
3、打印最终报告
printHeader()
defer printFooter()
4、关闭数据库链接
// open a database connection
defer disconnectFromDB()
*/
