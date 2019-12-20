package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig)
	// ToLower转小写
	// strings.ToLower(s) string
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	// ToUpper转大写
	// strings.ToUpper(s) string
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)
}
