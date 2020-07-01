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

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q} //slice、interface
	fmt.Println("Looping through shapes for area ...")
	fmt.Println("Looping through shapes for area is:", shapes)
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		//shapes[0]==>{5,3}调用的是Rectangle的接口方法 Area
		//shapes[1]===>&{5},调用的是Square的接口方法 Area
		//这是 多态 的 Go版本
		//多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法
		//或者说：同一种类型在不同的实例上似乎表现出不同的行为
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

/**
 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法。

 type ReadWrite interface {
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
