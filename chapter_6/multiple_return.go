package main

import "fmt"

var num int = 10
var numx2, numx3 int

/**
* 警告：
   return 或 return var 都是可以的。
   不过 return var = expression（表达式） 会引发一个编译错误：syntax error: unexpected =, expecting semicolon or newline or }。
   即使函数使用了命名返回值，你依旧可以无视它而返回明确的值。

   任何一个非命名返回值（使用非命名返回值是很糟的编程习惯）在 return 语句里面都要明确指出包含返回值的变量或是一个可计算的值（就像上面警告所指出的那样）。
* @Author Liumm
* @Date   2019-08-19
* @return {[type]}   [description]警告：

return 或 return var 都是可以的。
不过 return var = expression（表达式） 会引发一个编译错误：syntax error: unexpected =, expecting semicolon or newline or }。
即使函数使用了命名返回值，你依旧可以无视它而返回明确的值。

任何一个非命名返回值（使用非命名返回值是很糟的编程习惯）在 return 语句里面都要明确指出包含返回值的变量或是一个可计算的值（就像上面警告所指出的那样）。
*/
func main() {
    numx2, numx3 = getX2AndX3(num)
    PrintValues()
    numx2, numx3 = getX2AndX3_2(num)
    PrintValues()
}

func PrintValues() {
    fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) { //非命名返回值，没有声明变量放回名
    return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) { //命名返回值,声明了返回变量名
    x2 = 2 * input
    x3 = 3 * input
    // return x2, x3
    return
}
