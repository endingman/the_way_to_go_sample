package trans

import (
	"math"
)

var Pi float64

/**
变量除了可以在全局声明中初始化，也可以在init函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。

每个源文件都只能包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。

一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。
* @Author Liumm
* @Date   2019-12-17
* @return {[type]}   [description]
*/
func init() {
	// 使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
	// 可以通过 &i 来获取变量 i 的内存地址（第 4.9 节），例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中
	Pi = 4 * math.Atan(1) // init() function computes Pi
}
