package main

import "fmt"

/**
Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建： var arr1 = new([5]int)
* @Author Liumm
* @Date   2019-12-20
* @return {[type]}   [description]
*/
func main() {
	// 数组声明 var identifier [len]type
	var arr1 [5]int
	// 通过 for 初始化数组项
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	/**
	  在这里 i 也是数组的索引。当然这两种 for 结构对于切片（slices）（参考 第 7 章）来说也同样适用。
	  for i,_:= range arr1 {
	  ...
	  } //i 数组索引，arr1数组，range
	*/

	// 通过 for 打印数组元素
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}

	a := [...]string{"a", "b", "c", "d"}

	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

}
