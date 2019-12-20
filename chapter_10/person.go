package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

/**
 * 方法有一个类型为 *Person 的参数（因此对象本身是可以被改变的），以及三种调用这个方法的不同方式
 * @Author Liumm
 * @Date   2019-12-20
 * @param  {[type]}   p *Person       [description]
 * @return {[type]}     [description]
 */
func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1-struct as a value type:
	var pers1 Person // pers1是结构体类型变量
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"} // pers3是指向一个结构体类型变量的指针
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}

/**
 * type Interval struct {
    start int
    end   int
}
初始化方式：

intr := Interval{0, 3}            (A)
intr := Interval{end:5, start:1}  (B)
intr := Interval{end:5}           (C)
在（A）中，值必须以字段在结构体定义时的顺序给出，& 不是必须的。（B）显示了另一种方式，字段名加一个冒号放在值的前面，这种情况下值的顺序不必一致，并且某些字段还可以被忽略掉，就像（C）中那样。

结构体类型和字段的命名遵循可见性规则（第 4.2 节），一个导出的结构体类型中有些字段是导出的，另一些不是，这是可能的。

初始化一个结构体实例（一个结构体字面量：struct-literal）的更简短和惯用的方式如下：
    ms := &struct1{10, 15.5, "Chris"}
    // 此时ms的类型是 *struct1
*/
