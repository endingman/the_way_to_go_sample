package main

import (
    "fmt"
)

func main() {
    sum, mult, reduce := getX2AndX3(1, 2)
    fmt.Println("sum = %d, 2x mult = %d, 3x reduce = %d\n", sum, mult, reduce)
    sum, mult, reduce = getX2AndX3_2(1, 2)
    fmt.Println("sum = %d, 2x mult = %d, 3x reduce = %d\n", sum, mult, reduce)
}

func getX2AndX3(input1 int, input2 int) (int, int, int) {
    return input1 + input2, input1 * input2, input1 - input2
}

func getX2AndX3_2(input1 int, input2 int) (sum int, mult int, reduce int) {
    sum = input1 + input2
    mult = input1 * input2
    reduce = input1 - input2
    return
}
