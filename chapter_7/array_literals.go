package main

import "fmt"

func main() {
	// var arrAge = [5]int{18, 20, 15, 22, 16}

	// 数组长度同样可以写成 ... 或者直接忽略
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}

	// key: value syntax 只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
