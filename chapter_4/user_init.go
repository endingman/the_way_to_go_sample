package main

import (
	"fmt"

	"go_module/chapter_4/trans" //"goModule/test"  // 模块名 + 路径,go module模式引入自定义包方式
)

var twoPi = 2 * trans.Pi

func main() {
	// func Printf(format string, list of variables to be printed)
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
