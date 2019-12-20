// panic_defer.go
package main

import "fmt"

func main() {
	f()
	// 4、
	fmt.Println("Returned normally from f.")
}

func f() {
	// 3、
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	// 1
	fmt.Println("Calling g.")
	g(0) //2、g(0),g(1)g(2)g(3)g(4)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		//2、 g(4)
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	// 3、defer
	defer fmt.Println("Defer in g", i)
	//1、g(0),g(1)g(2)g(3)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
