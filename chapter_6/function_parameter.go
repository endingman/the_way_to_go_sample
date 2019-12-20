package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

/**
 * 函数做为参数
 * @Author Liumm
 * @Date   2019-12-20
 * @param  {[type]}   y int           [description]
 * @param  {[type]}   f func(int,     int)          [description]
 * @return {Function}   [description]
 */
func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}
