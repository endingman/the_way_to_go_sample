package main

import "fmt"

/**
* 结构体定义
* type identifier struct {
   field1 type1
   field2 type2
   ...}
*/
type struct1 struct {
  i1  int
  f1  float32
  str string
}

func main() {
  ms := new(struct1)
  // 像在面向对象语言所作的那样，可以使用点号符给字段赋值：structname.fieldname = value
  ms.i1 = 10
  ms.f1 = 15.5
  ms.str = "Chris"

  // 使用点号符可以获取结构体字段的值：structname.fieldname
  fmt.Printf("The int is: %d\n", ms.i1)
  fmt.Printf("The float is: %f\n", ms.f1)
  fmt.Printf("The string is: %s\n", ms.str)
  fmt.Println(ms)
}
