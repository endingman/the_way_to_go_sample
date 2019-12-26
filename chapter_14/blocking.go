// blocking.go
// throw: all goroutines are asleep - deadlock!
package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

/**
通信是一种同步形式：通过通道，两个协程在通信（协程会和）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。
甚至可以在通道两端互相阻塞对方，形成了叫做死锁的状态。Go 运行时会检查并 panic，停止程序。死锁几乎完全是由糟糕的设计导致的。
无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道。
 * @Author Liumm
 * @Date   2019-12-26
 * @return {[type]}   [description]
*/
func main() {
	// 这个程序会导致 panic：所有的协程都休眠了 - 死锁！
	// 因为没有接收者
	out := make(chan int)
	// 1、可以使用缓冲解决
	//out := make(chan int, 1) // solution 2

	//2、
	// go f1(out)  // solution 1
	out <- 2

	go f1(out)
}
