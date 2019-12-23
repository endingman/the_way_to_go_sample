package main

import (
	"fmt"

	"go_module/chapter_10/person"
)

/**
 考虑 person2.go 中的 person 包：类型 Person 被明确的导出了，但是它的字段没有被导出。例如在 use_person2.go 中 p.firstName 就是错误的。该如何在另一个程序中修改或者只是读取一个 Person 的名字呢？

这可以通过面向对象语言一个众所周知的技术来完成：提供 getter 和 setter 方法。对于 setter 方法使用 Set 前缀，对于 getter 方法只使用成员名。
 * @Author Liumm
 * @Date   2019-12-23
 * @return {[type]}   [description]
*/
func main() {
	p := new(person.Person)
	// p.firstName undefined
	// (cannot refer to unexported field or method firstName)
	// p.firstName = "Eric"
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName()) // Output: Eric
}
