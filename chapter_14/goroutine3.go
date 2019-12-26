package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	//通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。
	close(ch)
}

/**
	 如何检测到通道是否关闭或阻塞？

	第一个可以通过函数 close(ch) 来完成：这个将通道标记为无法通过发送操作 <- 接受更多的值；给已经关闭的通道发送或者再次关闭都会导致运行时的 panic。在创建一个通道后使用 defer 语句是个不错的办法（类似这种情况）：

	ch := make(chan float64)
	defer close(ch)
	第二个问题可以使用逗号，ok 操作符：用来检测通道是否被关闭。

	如何来检测可以收到没有被阻塞（或者通道没有被关闭）？

	v, ok := <-ch   // ok is true if v received value
	通常和 if 语句一起使用：

	if v, ok := <-ch; ok {
	  process(v)
	}
	或者在 for 循环中接收的时候，当关闭或者阻塞的时候使用 break：

	v, ok := <-ch
	if !ok {
	  break
	}
	process(v)
 * @Author Liumm
 * @Date   2019-12-26
 * @param  {[type]}   ch chan          string [description]
 * @return {[type]}      [description]
*/
func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}
