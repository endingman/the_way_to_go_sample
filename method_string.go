package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

/**
 * 当定义了一个有很多方法的类型时，十之八九你会使用 String() 方法来定制类型的字符串形式的输出，换句话说：一种可阅读性和打印性的输出。如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。
 * @Author Liumm
 * @Date   2019-08-21
 * @return {[type]}   [description]
 */
func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
