package main

import "fmt"

func main() {
    slice1 := make([]int, 0, 10)
    // load the slice, cap(slice1) is 10:
    for i := 0; i < cap(slice1); i++ {
        slice1 = slice1[0 : i+1]
        slice1[i] = i
        fmt.Printf("The length of slice is %d\n", len(slice1))
    }

    // print the slice:
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    }

    slice := make([]int, 10)
    slice2 := slice[10:10]
    fmt.Printf("The length of slice2 is %d\n", len(slice2))

    // slice3 := slice[10:11]
    // fmt.Printf("The length of slice3 is %d\n", len(slice3))
}
