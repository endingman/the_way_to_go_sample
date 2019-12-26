package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	  suck它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程(pump)必须写入 ch（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭。
	*/
	suck(pump())
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		/**
		  for 循环的 range 语句可以用在通道 ch 上，便可以从通道中获取值，像这样：

		  for v := range ch {
		      fmt.Printf("The value is %v\n", v)
		  }
		*/
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
