// 从控制台读取输入:
package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

/**
Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。Scanf 与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。
Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，与 Scanf 相同。
如果这些函数读取到的结果与您预想的不同，您可以检查成功读入数据的个数和返回的错误。
* @Author Liumm
* @Date   2019-12-24
* @return {[type]}   [description]
*/
func main() {
	fmt.Println("Please enter your full name: ")
	// Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}
