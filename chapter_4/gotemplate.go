/**
 Go 程序的一般结构:

 所有的结构将在这一章或接下来的章节中进一步地解释说明，但总体思路如下：

 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。

 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
 如果当前包是 main 包，则定义 main 函数。

然后定义其余的函数，首先是类型的方法，接着是按照main函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。
*/
package main

import (
	"fmt"
)

/**
如果你有多个类型需要定义，可以使用因式分解关键字的方式，例如：

type (
   IZ int
   FZ float64
   STR string
)
*/
const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package

}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	/**
	   类型转换：
	   valueOfTypeB = typeB(valueOfTypeA)
	   类型 B 的值 = 类型 B (类型 A 的值)

	  示例：

	  a := 5.0
	  b := int(a)

	  注：这段跟这个实例go程序结构无关，只是增加一个本篇节（4.2.6）的点在这里
	*/
}

func Func1() { // exported function Func1
	//...
}
