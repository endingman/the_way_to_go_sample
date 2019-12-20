package main

import "fmt"
import "time"

func main() {
	c := make(chan int)

	/**
	   一个无缓冲通道只能包含 1 个元素，有时显得很局限。我们给通道提供了一个缓存，可以在扩展的 make 命令中设置它的容量，如下：

	  buf := 100
	  ch1 := make(chan string, buf)
	  buf 是通道可以同时容纳的元素（这里是 string）个数

	  在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。

	  缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的 cap 函数可以返回缓冲区的容量。

	  如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是 0 或者未设置，通信仅在收发双方准备好的情况下才可以成功。

	  同步：ch :=make(chan type, value)
	*/
	// c := make(chan int, 50)

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
(15 s later):
received 10
sent 10
*/

/**
 协程通过在通道 ch 中放置一个值来处理结束的信号。main 协程等待 <-ch 直到从中获取到值。

我们期望从这个通道中获取返回的结果，像这样：

func compute(ch chan int){
    ch <- someComputation() // when it completes, signal on the channel.
}

func main(){
    ch := make(chan int)    // allocate a channel.
    go compute(ch)      // stat something in a goroutines
    doSomethingElseForAWhile()
    result := <- ch
}
*/
