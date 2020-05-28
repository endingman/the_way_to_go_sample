package main

import (
	"fmt"
	"os"
	"runtime"
)

/**
 * [main description]
 * @Author Liumm
 * @Date   2020-05-28
 * @return {[type]}   [description]
 */
// 系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。记住，所有的内存在 Go 中都是经过初始化的
func main() {
	// 变量声明 -> var identifier type
	/**多个可以这样 ->
	            var (
	                identifier1 type1,
	                identifier2 type2... )
	         example:
	            var (
	            a = 15
	            b = false
	            str = "Go says hello to the world!"
	            numShips = 50
	            city string
	        )

	   简短形式为：

	  同一类型的多个变量可以声明在同一行，如：

	  var a, b, c int
	  (这是将类型写在标识符后面的一个重要原因)

	  多变量可以在同一行进行赋值，如：

	  a, b, c = 5, 7, "abc"

	**/
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
