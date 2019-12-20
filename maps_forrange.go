package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	// map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序 (这里测试得到的结果是有序的)
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
	// map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序 (这里测试得到的结果是有序的)
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"} //给定{}初始化，成对初始化
	for key := range capitals {
		fmt.Println("Map item: Capital of key is:", key, "and value is", capitals[key])
	}
}
