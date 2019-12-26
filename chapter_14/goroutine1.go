package main

import (
	"fmt"
	"time"
)

/**
 main()，longWait() 和 shortWait() 三个函数作为独立的处理单元按顺序启动，然后开始并行运行
 Go 协程意味着并行（或者可以以并行的方式部署），协程一般来说不是这样的
 Go 协程通过通道来通信；协程通过让出和恢复操作来通信
 Go 协程比协程更强大，也很容易从协程的逻辑复用到 Go 协程。

 一个应用程序是运行在机器上的一个进程；进程是一个运行在自己内存地址空间里的独立执行体。一个进程由一个或多个操作系统线程组成，这些线程其实是共享同一个内存地址空间的一起工作的执行体。几乎所有 ' 正式 ' 的程序都是多线程的，以便让用户或计算机不必等待，或者能够同时服务多个请求（如 Web 服务器），或增加性能和吞吐量（例如，通过对不同的数据集并行执行代码）。一个并发程序可以在一个处理器或者内核上使用多个线程来执行任务，但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行。

并行是一种通过使用多处理器以提高速度的能力。所以并发程序可以是并行的，也可以不是。

公认的，使用多线程的应用难以做到准确，最主要的问题是内存中的数据共享，它们会被多线程以无法预知的方式进行操作，导致一些无法重现或者随机的结果（称作 竞态）。
* @Author Liumm
* @Date   2019-09-06
* @return {[type]}   [description]
*/
func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9) //纳秒（ns，符号 1e9 表示 1 乘 10 的 9 次方，e = 指数）
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
