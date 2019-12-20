package main

import "fmt"

/**
 * map 是引用类型，可以使用如下声明：
 * var map1 map[keytype]valuetype
 * var map1 map[string]int
 * @Author Liumm
 * @Date   2019-08-20
 * @return {[type]}   [description]
 */
func main() {
	var mapLit map[string]int
	// var mapCreated map[string]float32
	// map 的初始化：var map1 = make(map[keytype]valuetype)。
	// 或者简写为：map1 := make(map[keytype]valuetype)
	var mapAssigned map[string]int
	// map literals 的使用方法： map 可以用 {key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样
	mapLit = map[string]int{"one": 1, "two": 2}
	/**
	 * map 是 引用类型 的： 内存用 make 方法来分配。
	 * map 的初始化：var map1 = make(map[keytype]valuetype)
	 * @type {[type]}
	 */
	// key 可以是任意可以用 == 或者！= 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key
	// * 注意 如果你错误的使用 new () 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址
	/**
	mapCreated := new(map[string]float32)
	接下来当我们调用：mapCreated["key1"] = 4.5 的时候，编译器会报错：

	invalid operation: mapCreated["key1"] (index of type *map[string]float32).
	*/
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}
