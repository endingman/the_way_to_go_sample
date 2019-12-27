package main

import "fmt"

/**
* 结构体定义
* type identifier struct {
   field1 type1
   field2 type2
   ...}

   - 切片、映射和通道，使用 make
   - 数组、结构体和所有的值类型，使用 new

  结构体里的字段都有 名字，像 field1、field2 等，如果字段在代码中从来也不会被用到，那么可以命名它为 _。

  结构体的字段可以是任何类型，甚至是结构体本身（参考第 10.5 节），也可以是函数或者接口（参考第 11 章）。可以声明结构体类型的一个变量，然后像下面这样给它的字段赋值：
  var s T
  s.a = 5
  s.b = 8
*/
type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	/**
	  使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)，如果需要可以把这条语句放在不同的行（比如定义是包范围的，但是分配却没有必要在开始就做）
	  * @type {[type]}
	*/
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
