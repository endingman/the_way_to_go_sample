package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	time 包中有一些有趣的功能可以和通道组合使用。

	其中就包含了 time.Ticker 结构体，这个对象以指定的时间间隔重复的向通道 C 发送时间值：

	type Ticker struct {
	    C <-chan Time // the channel on which the ticks are delivered.
	    // contains filtered or unexported fields
	    ...
	}
	时间间隔的单位是 ns（纳秒，int64），在工厂函数 time.NewTicker 中以 Duration 类型的参数传入：func Newticker(dur) *Ticker。
	*/
	tick := time.Tick(1e8)

	/**
	func After(d Duration) <-chan Time
	在 Duration d 之后，当前时间被发到返回的通道；所以它和 NewTimer(d).C 是等价的；它类似 Tick()，但是 After() 只发送一次时间
	*/
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
