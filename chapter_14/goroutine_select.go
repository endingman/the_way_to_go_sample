package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		/**
		  从不同的并发执行的协程中获取值可以通过关键字 select 来完成，它和 switch 控制语句非常相似（章节 5.3）也被称作通信开关；它的行为像是 “你准备好了吗” 的轮询机制；select 监听进入通道的数据，也可以是用通道发送值的时候。

		   select {
		   case u:= <- ch1:
		           ...
		   case v:= <- ch2:
		           ...
		           ...
		   default: // no value ready to be received
		           ...
		   }
		   default 语句是可选的；fallthrough 行为，和普通的 switch 相似，是不允许的。在任何一个 case 中执行 break 或者 return，select 就结束了。

		   select 做的就是：选择处理列出的多个通信情况中的一个。

		   如果都阻塞了，会等待直到其中一个可以处理
		   如果多个可以处理，随机选择一个
		   如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。
		   在 select 中使用发送操作并且有 default 可以确保发送不被阻塞！如果没有 case，select 就会一直阻塞。

		   select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出。
		*/
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
