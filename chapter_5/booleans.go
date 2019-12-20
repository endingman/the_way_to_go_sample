package main

import (
	"fmt"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

// func init() {
//     if runtime.GOOS == "windows" {
//         prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
//     } else { //Unix-like
//         prompt = fmt.Sprintf(prompt, "Ctrl+D")
//     }
// }

func main() {
	bool1 := true
	if bool1 {
		fmt.Println("The value is true\n")
	} else {
		fmt.Println("The value is false\n")
	}

	x := Abs(1)
	fmt.Println(x)

	y := 2
	fmt.Println(isGreater(x, y))
}

/**
if 可以包含一个初始化语句（如：给一个变量赋值）。这种写法具有固定的格式（在初始化语句后方必须加上分号）：

if initialization; condition {
    // do something
}
*/

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false

	// ====> return x>y
}
