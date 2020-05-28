package main

var a string

func main() {
	// 全局的a
	a = "G"
	print(a) //G
	f1()
}

func f1() {
	// 函数里的a
	a := "O"
	print(a) //O
	f2()
}

func f2() {
	print(a) //main赋值的全局G
}

// result ->GOG
