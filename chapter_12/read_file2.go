package main

import (
	"fmt"
	"os"
)

func main() {
	// os.open打开文件
	file, err := os.Open("products2.txt")
	if err != nil {
		panic(err)
	}
	// defer 在return前运行,类似栈,先进后出
	defer file.Close()

	var col1, col2, col3 []string //slice

	for {
		var v1, v2, v3 string
		// 如果数据是按列排列并用空格分隔的，可以使用 fmt 包提供的以 FScan 开头的一系列函数来读取他们
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
