// panic_package.go
package main

import (
	"fmt"
	"go_module/parse" //"goModule/test"  // 模块名 + 路径,go module模式引入自定义包方式
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n  ", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err) // here String() method from ParseError is used
			continue
		}
		fmt.Println(nums)
	}
}
