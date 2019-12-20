package main

func main() {
	// 一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）
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
