package main

import (
	"fmt"

	"strconv"
)

type Person struct {
	Name string

	salary float64

	chF chan func()
}

func NewPerson(name string, salary float64) *Person {

	p := &Person{name, salary, make(chan func())}

	go p.backend()

	return p

}

func (p *Person) backend() {

	for f := range p.chF {

		f()

	}

}

// 设置 salary.

func (p *Person) SetSalary(sal float64) {

	p.chF <- func() {
		p.salary = sal
	}

}

// 取回 salary.

func (p *Person) Salary() float64 {

	fChan := make(chan float64)

	p.chF <- func() {
		fChan <- p.salary
	}

	return <-fChan

}

func (p *Person) String() string {

	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)

}

/**
 为了保护一个对象的并发修改，我们可以使用一个后台的协程来顺序执行一个匿名函数，而不是通过同步 互斥锁（Mutex） 进行锁定。

在下面的程序中，我们有一个 Person 类型，它包含了一个匿名函数类型的通道字段 chF。它在构造器方法 NewPerson 中初始化，用一个协程启动一个 backend() 方法。这个方法在一个无限 for 循环中执行所有被放到 chF 上的函数，有效的序列化他们，从而提供安全的并发访问。改变和获取 salary 可以通过一个放在 chF 上的匿名函数来实现，backend() 会顺序执行它们。注意如何在 Salary 方法中的闭合（匿名）函数中去包含 fChan 通道。

这是一个简化的例子，并且它不应该在这种情况下应用，但是它展示了如何在更复杂的情况下解决问题。

 * @Author Liumm
 * @Date   2019-12-26
 * @return {[type]}   [description]
*/
func main() {

	bs := NewPerson("Smith Bill", 2500.5)

	fmt.Println(bs)

	bs.SetSalary(4000.25)

	fmt.Println("Salary changed:")

	fmt.Println(bs)

}
