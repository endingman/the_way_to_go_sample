/**
 * 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

通过如下格式定义接口：

type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}

类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

实现某个接口的类型（除了实现接口方法外）可以有其他的方法。

一个类型可以实现多个接口。

接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

*/
package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

/**
在 main() 方法中创建了一个 Square 的实例。
在主程序外边定义了一个接收者类型是 Square 方法的 Area()，用来计算正方形的面积：结构体 Square 实现了接口 Shaper
 * @Author Liumm
 * @Date   2019-12-23
 * @return {[type]}   [description]
*/
func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

/**
 上面的程序定义了一个结构体 Square 和一个接口 Shaper，接口有一个方法 Area()。

在 main() 方法中创建了一个 Square 的实例。在主程序外边定义了一个接收者类型是 Square 方法的 Area()，用来计算正方形的面积：结构体 Square 实现了接口 Shaper 。

所以可以将一个 Square 类型的变量赋值给一个接口类型的变量：areaIntf = sq1 。

现在接口变量包含一个指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()。当然也可以直接在 Square 的实例上调用此方法，但是在接口实例上调用此方法更令人兴奋，它使此方法更具有一般性。接口变量里包含了接收者实例的值和指向对应方法表的指针。

这是 多态 的 Go 版本，多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法，或者说：同一种类型在不同的实例上似乎表现出不同的行为。
*/
