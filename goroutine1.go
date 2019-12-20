package main

import (
	"fmt"
	"time"
)

/**
 * main()，longWait() 和 shortWait() 三个函数作为独立的处理单元按顺序启动，然后开始并行运行
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
	time.Sleep(10 * 1e9)
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
