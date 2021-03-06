package main

import "fmt"

func main() {
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	//int16 不能够被隐式转换为 int32。
	//通过显式转换来避免这个问题
	m = int32(n)
	/**
	  格式化字符串里，%d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数字）
	  %g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法），
	  %0d 用于规定输出定长的整数，其中开头的数字 0 是必须的。

	  %n.mg 用于表示数字 n 并精确到小数点后 m 位，
	  除了使用 g 之外，还可以使用 e 或者 f，例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00
	         * @type {[type]}
	*/
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}
