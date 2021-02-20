/**
你可以通过使用包的别名来解决包名之间的名称冲突，
或者说根据你的个人喜好对包名进行重新设置，如：import fm "fmt"。
下面的代码展示了如何使用包的别名：
 */

package main

import fm "fmt"

func main()  {
	fm.Println("Hello World!")
}