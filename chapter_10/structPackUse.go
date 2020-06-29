package main

import (
	"fmt"
	structPack "go_module/chapter_10/struct_pack" //"goModule/test"  // 模块名 + 路径,go module模式引入自定义包方式
)

func main() {
	// 使用自定义包中的结构体
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}
