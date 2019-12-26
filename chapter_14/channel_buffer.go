package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 50)

	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}

/* Output:
sending 10
sent 10   // prints immediately
no further output, because main() then stops
需要使用信号量模式才会输出

也可以使用通道来达到同步的目的，这个很有效的用法在传统计算机中称为信号量（semaphore）。或者换个方式：通过通道发送信号告知处理已经完成（在协程中）。

在其他协程运行时让 main 程序无限阻塞的通常做法是在 main 函数的最后放置一个 {}。

也可以使用通道让 main 程序等待协程完成，就是所谓的信号量模式
*/
