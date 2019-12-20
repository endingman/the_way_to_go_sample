package main

import "fmt"

func main() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	/**
	   如果你只需要索引，你可以忽略第二个变量，例如：

	  for ix := range seasons {
	      fmt.Printf("%d", ix)
	  }
	       * @type {[type]}
	*/
	// _ 可以用于忽略索引
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
