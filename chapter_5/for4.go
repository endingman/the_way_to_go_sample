package main

func main() {
	// 下面的示例中包含了嵌套的循环体（for4.go），break 只会退出最内层的循环：
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			print(j)
		}
		print("  ")
	}
}
