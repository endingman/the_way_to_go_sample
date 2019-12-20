package main

import (
  "fmt"
  "strconv"
)

func main() {
  var orig string = "ABC"
  // var an int
  var newS string
  // var err error

  fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
  // anInt, err = strconv.Atoi(origStr)
  an, err := strconv.Atoi(orig)
  if err != nil {
    fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
    return
  }

  // ==>error习惯用法
  /**
        value, err := pack1.Function1(param1)
    if err != nil {
        fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
        return err
    }
    // 未发生错误，继续执行：
  */

  fmt.Printf("The integer is %d\n", an)
  an = an + 5
  newS = strconv.Itoa(an)
  fmt.Printf("The new string is: %s\n", newS)
}
