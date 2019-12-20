package main

import (
	"fmt"
	"strconv"
)

/**
 * 字符串相关的类型转换都是通过 strconv 包实现的
 * @Author Liumm
 * @Date   2019-12-17
 * @return {[type]}   [description]
 */
func main() {
	var orig string = "666"
	var an int
	var newS string

	// 用于获取程序运行的操作系统平台下 int 类型所占的位数，如：strconv.IntSize
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	// strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)

	an = an + 5
	//// strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数
	newS = strconv.Itoa(an)

	fmt.Printf("The new string is: %s\n", newS)
}
