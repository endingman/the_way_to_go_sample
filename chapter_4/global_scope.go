package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O"//这里区别于local_scope并没有重新使用短声明只是把O赋值全局a
	print(a)
}

// GOO
