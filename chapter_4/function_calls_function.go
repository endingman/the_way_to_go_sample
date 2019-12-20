package main

var a string

func main() {
	// 全局的a
	a = "G"
	print(a)
	f1()
}

func f1() {
	// 函数里的a
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
