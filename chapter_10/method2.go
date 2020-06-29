package main

import "fmt"

type IntVector []int

/**
 * 非结构体类型上方法的例子：
 * 接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。
 * 但是接收者不能是一个接口类型（参考 第 11 章），因为接口是一个抽象定义，但是方法却是具体实现；
 * @Author Liumm
 * @Date   2019-12-23
 * @param  {[type]}   v IntVector)    Sum() (s int [description]
 * @return {[type]}     [description]
 */
func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum()) // 输出是6
}

/**
函数将变量作为参数：Function1(recv) => 需要申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体：
func flushICache(begin, end uintptr) // implemented externally


方法在变量上被调用：recv.Method1()=>  func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... },如果方法不需要使用 recv 的值，可以用 _ 替换它，比如：func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }

在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。

不要忘记 Method1 后边的括号 ()，否则会引发编译器错误：method recv.Method1 is not an expression, must be called

接收者必须有一个显式的名字，这个名字必须在方法中被使用。

receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。

在 Go 中，（接收者）类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接收者来建立。

方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。
*/
