// Learn more or give us feedback
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "fmt"

/**
一个请求结构体类似如下形式，内嵌了一个回复 channel：
*/
type Request struct {
	a, b   int
	replyc chan int // 请求中的回复频道
}

/**
服务端可以在一个 goroutine 里面为每个请求都分配一个 run () 函数，这个函数会把 binOp 类型的操作作用于整数，然后通过回复 channel 发送结果：
*/
type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

/**
服务端通过死循环来从 chan *Request 接收请求，为了避免长时间运行而导致阻塞，可以为每个请求都开一个 goroutine 来处理
*/
func server(op binOp, service chan *Request) {
	for {
		req := <-service // 请求到达这里
		// 开启请求的 Goroutine ：
		go run(op, req) // 不要等待 op
	}
}

/**
使用 startServer 函数来启动服务的自有的协程（goroutine）
*/
func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request)
	go server(op, reqChan)
	return reqChan
}

/**
 Client-server 类的应用是协程（goroutine）和频道（channel）的大显身手的闪光点。

客户端可以是任何一种运行在任何设备上的，且需要来自服务端信息的一种程序，所以它需要发送请求。 服务端接收请求，做一些处理，然后把给客户端发送响应信息。在通常情况下，就是多个客户端（很多请求）对一个（或几个）服务端。一个常见例子就是我们使用的发送网页请求的客户端浏览器。然后一个 web 服务器将响应网页发回给浏览器。

在 Go 中，服务端通常会在一个协程（goroutine）里操作对一个客户端的响应，所以协程和客户端请求是一一对应的。一种典型的做法就是客户端请求本身包含了一个频道（channel），服务端可以用它来发送响应。
 * @Author Liumm使用 startServer 函数来启动服务的自有的协程（goroutine）
 * @Date   2019-12-26
 * @return {[type]}   [description]
*/
func main() {
	adder := startServer(func(a, b int) int {
		return a + b
	})
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}
	// 校验：
	for i := N - 1; i >= 0; i-- { // 顺序无所谓
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}
	fmt.Println("done")
}
