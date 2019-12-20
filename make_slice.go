package main

import "fmt"

/**
 * 当相关数组还没有定义时，我们可以使用 make () 函数来创建一个切片 同时创建好相关数组：var slice1 []type = make([]type, len)。

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
