package main

import (
    "fmt"
    "go_module/struct_pack" //"goModule/test"  // 模块名 + 路径,go module模式引入自定义包方式
)

func main() {
    struct1 := new(structPack.ExpStruct)
    struct1.Mi1 = 10
    struct1.Mf1 = 16

    fmt.Printf("Mi1 = %d\n", struct1.Mi1)
    fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}