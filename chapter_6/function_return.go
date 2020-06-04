package main

import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		fmt.Printf("The Add2 b is: %v\n", b)
		return b + 2 //5
	}
}

/**
 * 应用闭包：将函数作为返回值
 */
func Adder(a int) func(b int) int {
	fmt.Printf("The a is: %v\n", a)
	return func(b int) int {
		fmt.Printf("The b is: %v\n", b)
		return a + b //5
	}
}
