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
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		//shapes[0]==>{5,3}调用的是Rectangle结构体的方法Area，shapes[1]===>&{5},调用的是Square的接口方法Area
		//多态行为
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
