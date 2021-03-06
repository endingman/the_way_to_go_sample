package main

import "fmt"

/**
 * 切片是一个引用类型
 * 所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用
 * @Author Liumm
 * @Date   2019-12-20
 * @return {[type]}   [description]
 */
func main() {

	/**
	  var slice1 []type = arr1[:] 那么 slice1 就等于完整的 arr1 数组（所以这种表示方式是 arr1[0:len(arr1)] 的一种缩写）。另外一种表述方式是：slice1 = &arr1。

	  arr1[2:] 和 arr1[2:len(arr1)] 相同，都包含了数组从第三个到最后的所有元素。

	  arr1[:3] 和 arr1[0:3] 相同，包含了从第一个到第三个元素（不包括第三个，就是包含开始的第几个不含结束的那个）
	*/

	var arr1 [6]int              //数组指定长度，值没有初始化，使用0填充
	var slice1 []int = arr1[2:5] // item at index 5 not included! slice [strat:end] ===> start and end 是第几个而非下标,这里是第二个到第五个不包括第5个

	// load the array with integers: 0,1,2,3,4,5
	//  切片 s 来说该不等式永远成立：0 <= len(s) <= cap(s)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
}

/**
 * 如果 s2 是一个 slice，你可以将 s2 向后移动一位 s2 = s2[1:]，但是末尾没有移动。切片只能向后移动，s2 = s2[-1:] 会导致编译错误。切片不能被重新分片以获取数组的前一个元素。

注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针！！
*/

/**
 * 如果你有一个函数需要对数组做操作，你可能总是需要把参数声明为切片。当你调用该函数时，把数组分片，创建为一个 切片引用并传递给该函数。这里有一个计算数组元素和的方法:

func sum(a []int) int {
    s := 0
    for i := 0; i < len(a); i++ {
        s += a[i]
    }
    return s
}

func main() {
    var arr = [5]int{0, 1, 2, 3, 4}
    sum(arr[:])
}
*/
