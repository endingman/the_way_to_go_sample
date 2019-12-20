package main

/**
左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。

右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。
*/
const (
	AvailableMemory = 10 << 20 //(<<左移)

	// 10 MB, 示例

	AverageMemoryPerRequest = 10 << 10 //(左移)

	// 右移>>

	// 10 KB

	MAXREQS = AvailableMemory / AverageMemoryPerRequest
	// 原文中说 MAXREQS 是 1000，实际计算是 1024 ，后面按照原文的 1000 来描述

)

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b int

	replyc chan int
}

func process(r *Request) {

	// Do something 做任何事

	// 可能需要很长时间并使用大量内存或CPU

}

func handle(r *Request) {

	process(r)

	// 信号完成：开始启用下一个请求

	// 将 sem 的缓冲区释放一个位置

	<-sem

}

func Server(queue chan *Request) {

	for {

		sem <- 1

		// 当通道已满（1000 个请求被激活）的时候将被阻塞

		// 所以停在这里等待，直到 sem 有容量（被释放），才能继续去处理请求

		// (doesn’t matter what we put in it)

		request := <-queue

		go handle(request)

	}

}

func main() {

	queue := make(chan *Request)

	go Server(queue)

}
