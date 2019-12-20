package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)

	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25

	/**
	 * val1, isPresent = map1[key1]
	 * isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent 为true；
	 * 如果key1不存在，val1就是一个空值，并且 isPresent 会返回 false。
	 * @type {Boolean}
	 */
	value, isPresent = map1["Beijing"]

	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)

	// delete an item:
	/**
	 * 从 map1 中删除 key1：
	 * 直接 delete(map1, key1) 就可以。
	 * 如果 key1 不存在，该操作不会产生错误。
	 */
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}
}
