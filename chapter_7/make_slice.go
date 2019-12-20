package main

import "fmt"

/**
 * 值类型可以使用new，指针类型使用make，slice是指针类型，本身就是一个引用

 * 也可以简写为 slice1 := make([]type, len)，这里 len 是数组的长度并且也是 slice 的初始长度。

 * 所以定义 s2 := make([]int, 10)，那么 cap(s2) == len(s2) == 10。

 * make 接受 2 个参数：元素的类型以及切片的元素个数。

 * 如果你想创建一个 slice1，它不占用整个数组，而只是占用以 len 为个数个项，那么只要：slice1 := make([]type, len, cap)。

 * make 的使用方式是：func make([]T, len, cap)，其中 cap 是可选参数。
 * @Author Liumm
 * @Date   2019-08-20
 * @return {[type]}   [description]
 */
func main() {
	var slice1 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i //5*0,5*1...
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}

/**
 new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。

make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel（参见第 8 章，第 13 章）。

换言之，new 函数分配内存，make 函数初始化；
*/
