package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS string = "Hi there! "
	var newS string
	/**
	  Repeat 用于重复 count 次字符串 s 并返回一个新的字符串：即字符串重复几遍输出

	  strings.Repeat(s, count int) string
	       * @type {[type]}
	*/
	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}
