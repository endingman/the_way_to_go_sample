package main

import (
	"fmt"

	"go_module/chapter_9/pack1" //go module模式引入自定义包 模块名称/相对路径.../包名
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}
