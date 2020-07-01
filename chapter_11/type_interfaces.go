package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	/**
	   if v, ok := varI.(T); ok {  // checked type assertion  varI=>接口类型变量，T指针类型的变量,v =>value,ok=>true/false
	      Process(v)
	      return
	  }
	  // varI is not of type T
	  如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；否则 v 是类型 T 的零值，ok 是 false，也没有运行时错误发生。
	* @type {[type]}
	*/
	if t, ok := areaIntf.(*Square); ok { //areaIntf=>sq1=>Square,所以ok
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok { //areaIntf=>sq1=>Square,所以not ok
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
