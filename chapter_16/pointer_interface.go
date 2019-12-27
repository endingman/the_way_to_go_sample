package main

import (
	"fmt"
)

type nexter interface {
	next() byte
}

func nextFew1(n nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}

// 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
func nextFew2(n *nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next() // 编译错误:n.next未定义（*nexter类型没有next成员或next方法）
	}
	return b
}
func main() {
	fmt.Println("Hello World!")
}
