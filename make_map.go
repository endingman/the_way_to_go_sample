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
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	/**
	 * map 是 引用类型 的： 内存用 make 方法来分配。
	 * map 的初始化：var map1 = make(map[keytype]valuetype)
	 * @type {[type]}
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
