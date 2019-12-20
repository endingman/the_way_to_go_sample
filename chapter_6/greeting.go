package main

// 函数被调用的基本格式如下：
// pack1.Function(arg1, arg2, …, argn)

func main() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}

/**
函数可以将其他函数调用作为它的参数，只要这个被调用函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的，例如：

假设 f1 需要 3 个参数 f1(a, b, c int)，同时 f2 返回 3 个参数 f2(a, b int) (int, int, int)，就可以这样调用 f1：f1(f2(a, b))。
 * @Author Liumm
 * @Date   2019-12-18
 * @return {[type]}   [description]
*/
func greeting() {
	println("In greeting: Hi!!!!!")
}

/**
函数也可以以申明的方式被使用，作为一个函数类型，就像：
type binOp func(int, int) int
*/

/**
函数值（functions value）之间可以相互比较：如果它们引用的是相同的函数或者都是 nil 的话，则认为它们是相同的函数。函数不能在其它函数里面声明（不能嵌套），不过我们可以通过使用匿名函数（参考 第 6.8 节）来破除这个限制。
*/
