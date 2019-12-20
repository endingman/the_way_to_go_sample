package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

/**
类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
一个类型可以实现多个接口。
接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。
 * @Author Liumm
 * @Date   2019-08-22
 * @return {[type]}   [description]
*/
func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4} //初始化结构体以及interface
	//showValue参数是interface，最后运行调用对应结构体实现interface具体方法
	/**
	        func (s stockPosition) getValue() float32 {
	      return s.sharePrice * s.count
	  }
	*/

	showValue(o) //stockPosition->getvalue stockPosition实现的接口方法
	o = car{"BMW", "M3", 66500}
	showValue(o) //car->getvalue car实现的接口方法
}

/**
 * 接口嵌套：一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法
 * type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}

type Lock interface {
    Lock()
    Unlock()
}

type File interface {
    ReadWrite
    Lock
    Close()
}
*/
