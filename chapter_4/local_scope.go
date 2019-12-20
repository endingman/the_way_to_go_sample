package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	// 全局的a
	print(a)
}

func m() {
	// 函数块里定义，不是全局的a
	a := "O"
	print(a)
}
