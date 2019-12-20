package main

import "fmt"

/**
 * 把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
 * 传递数组的指针
 * 使用数组的切片
 * @Author Liumm
 * @Date   2019-08-20
 * @return {[type]}   [description]
 */
func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) { //命名返回值
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
