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
	a := "O" //这里重新使用了:=短声明，不再是全局外的a
	print(a)
}

// result -> GOG
