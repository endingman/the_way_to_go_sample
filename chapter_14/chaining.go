package main

import (
	"flag"

	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}

/**
chaining.go 再次演示了启动大量的协程是多么的容易。
它发生在 mian 函数的 for 循环中。
在循环之后，向 rightmost 通道中插入 0 ，在不到 1.5 s 的时间执行了 100000 个协程，并将结果 100000 打印。
这个程序还演示了如何通过命令行的参数定义一个协程的数量，并通过 flag.Int 解析，例如： chaining -n=7000 （编译后通过命令行执行），可以生成 7000 个协程
* @Author Liumm
* @Date   2019-09-10
* @return {[type]}   [description]
*/
func main() {

	flag.Parse()

	leftmost := make(chan int)

	var left, right chan int = nil, leftmost

	for i := 0; i < *ngoroutine; i++ {

		left, right = right, make(chan int)

		go f(left, right)

	}

	right <- 0

	// start the chaining 开始链接

	x := <-leftmost // wait for completion 等待完成

	fmt.Println(x)

	// 结果： 100000 ， 大约 1,5s （我实际测试只用了不到200ms）

}
