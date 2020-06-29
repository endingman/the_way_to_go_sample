package main

type Foo map[string]string //指针

type Bar struct { //值类型
	thingOne string
	thingTwo int
}

func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// NOT OK 、make的只有slice、map、chan这三个，这三个是指针类型，其他都是new
	z := make(Bar) // 编译错误：cannot make type Bar
	(*z).thingOne = "hello"
	(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK
	u := new(Foo)         //new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"
}
