package main

import "fmt"

/**
结构体就像是类的一种简化形式
*/
type TwoInts struct {
	a int
	b int
}

/**
在 Go 语言中，结构体就像是类的一种简化形式，那么面向对象程序员可能会问：类的方法在哪里呢？在 Go 中有一个概念，它和方法有着同样的名字，并且大体上意思相同：Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。

接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。但是接收者不能是一个接口类型（参考 第 11 章），因为接口是一个抽象定义，但是方法却是具体实现；如果这样做会引发一个编译错误：invalid receiver type…。

最后接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。

一个类型加上它的方法等价于面向对象中的一个类。一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。

类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集。

因为方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法。但是如果基于接收者类型，是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，比如在同一个包里这么做是允许的：

func (a *denseMatrix) Add(b Matrix) Matrix
func (a *sparseMatrix) Add(b Matrix) Matrix

func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
在方法名之前，func 关键字之后的括号中指定 receiver。

如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。

如果 recv 一个指针，Go 会自动解引用。

如果方法不需要使用 recv 的值，可以用 _ 替换它，比如：

func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
recv 就像是面向对象语言中的 this 或 self，但是 Go 中并没有这两个关键字。随个人喜好，你可以使用 this 或 self 作为 receiver 的名字。下面是一个结构体上的简单方法的例子
*/
func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

/**
 * 结构体的方法，类似于其他语言类中的方法
 */
func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

/**
 * 结构体的方法，类似于其他语言类中的方法
 */
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

/**
函数和方法的区别#

函数将变量作为参数：Function1(recv)

方法在变量上被调用：recv.Method1()  一般格式：func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。

不要忘记 Method1 后边的括号 ()，否则会引发编译器错误：method recv.Method1 is not an expression, must be called

接收者必须有一个显式的名字，这个名字必须在方法中被使用。

receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。

在 Go 中，（接收者）类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接收者来建立。

方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。
*/
