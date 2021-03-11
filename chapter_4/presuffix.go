package main

import (
	"fmt"
	"strings"
)

/**

HasPrefix 判断字符串 s 是否以 prefix 开头：
strings.HasPrefix(s, prefix string) bool

HasSuffix 判断字符串 s 是否以 suffix 结尾：
strings.HasSuffix(s, suffix string) bool

 * @Author Liumm
 * @Date   2019-12-17
 * @return {[type]}   [description]
*/
func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}

/**
Contains 判断字符串 s 是否包含 substr：
strings.Contains(s, substr string) bool

Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
strings.Index(s, str string) int

LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
strings.LastIndex(s, str string) int
*/
