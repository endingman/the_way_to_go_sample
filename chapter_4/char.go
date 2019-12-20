package main

import "fmt"

/**
字符串拼接符 +
两个字符串 s1 和 s2 可以通过 s := s1 + s2 拼接在一起。
* @Author Liumm
* @Date   2019-12-17
* @return {[type]}   [description]
*/
func main() {
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”
}
