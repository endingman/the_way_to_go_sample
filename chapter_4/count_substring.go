package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"
	/**
	  Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数：

	  strings.Count(s, str string) int
	       * @type {String}
	*/
	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}
