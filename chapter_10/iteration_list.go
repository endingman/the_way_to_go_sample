package main

import "container/list"

func (p *list.List) Iter() {
	// ...
}

/**
 类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float 或类似这些的类型上定义方法。试图在 int 类型上定义方法会得到一个编译错误：

cannot define new methods on non-local type int
比如想在 time.Time 上定义如下方法：

func (t time.Time) first3Chars() string {
    return time.LocalTime().String()[0:3]
}
类型在其他的，或是非本地的包里定义，在它上面定义方法都会得到和上面同样的错误。

但是有一个间接的方式：可以先定义该类型（比如：int或float）的别名类型，然后再为别名类型定义方法。或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
 * @Author Liumm
 * @Date   2019-12-23
 * @return {[type]}   [description]
*/
func main() {
	lst := new(list.List)
	for range lst.Iter() {
	}
}
