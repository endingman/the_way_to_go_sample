package main

import "fmt"

/**
函数中数组作为参数传入时，如 func1(arr2)，会产生一次数组拷贝，func1 方法不会修改原始的数组 arr2。
如果你想修改原数组，那么 arr2 必须通过 & 操作符以引用方式传过来
另一种方法就是生成数组切片并将其传递给函数
 * @Author Liumm
 * @Date   2019-08-20
 * @param  {[type]}   a [3]int        [description]
 * @return {[type]}     [description]
*/
func f(a [3]int) {
	a[1] = 1
	fmt.Println(a)
}

func fp(a *[3]int) {
	a[1] = 1
	fmt.Println(a)
}

func main() {
	var ar [3]int
	f(ar)           // passes a copy of ar 不会修改ar的值
	fmt.Println(ar) //ar没有被改变
	fp(&ar)         // passes a pointer to ar 可以修改ar的值
	fmt.Println(ar) //ar被改变
}
