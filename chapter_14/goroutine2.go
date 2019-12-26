package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// main() 函数中启动了两个协程：sendData() 通过通道 ch 发送了 5 个字符串，getData() 按顺序接收它们并打印出来。

	// 如果 2 个协程需要通信，你必须给他们同一个通道作为参数才行。
	go sendData(ch)
	go getData(ch)

	/**
	  尝试一下如果注释掉 time.Sleep(1e9) 会如何。

	  我们发现协程之间的同步非常重要：

	  main () 等待了 1 秒让两个协程完成，如果不这样，sendData () 就没有机会输出。
	  getData () 使用了无限循环：它随着 sendData () 的发送完成和 ch 变空也结束了。
	  如果我们移除一个或所有 go 关键字，程序无法运行，Go 运行时会抛出 panic(死锁)：
	  ---- Error run E:/Go/Goboek/code examples/chapter 14/goroutine2.exe with code Crashed ---- Program exited with code -2147483645: panic: all goroutines are asleep-deadlock!
	*/

	time.Sleep(1e9)
}

/**
 通信操作符 <-
 这个操作符直观的标示了数据的传输：信息按照箭头的方向流动。

 流向通道（发送）

 ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）

 从通道流出（接收），三种方式：

 int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；
 假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch。

 <- ch 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：

 if <- ch != 1000{
    ...
 }
* @Author Liumm
* @Date   2019-09-06
* @param  {[type]}   ch chan          string [description]
* @return {[type]}      [description]
*/
func sendData(ch chan string) {
	/**
	  流向通道（发送）
	  ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）
	*/
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		/**
		  从通道流出（接收），三种方式：
		  int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；
		  假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch。
		*/
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
