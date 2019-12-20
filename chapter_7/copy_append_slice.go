package main

import "fmt"

/**
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
下面的代码描述了从拷贝切片的 copy 函数和向切片追加新元素的 append 函数。

append 非常有用，它能够用于各种方面的操作
将切片 b 的元素追加到切片 a 之后：a = append(a, b...)

复制切片 a 的元素到新的切片 b 上：

b = make([]T, len(a))
copy(b, a)
删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)

切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)

为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)

在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)

在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)

在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)

取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1], a[:len(a)-1]

将元素 x 追加到切片 a：a = append(a, x)
* @Author Liumm
* @Date   2019-12-20
* @return {[type]}   [description]
*/
func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from) //把sl_from复制到sl_to对应位置
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6) //追加   ===> append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；追加的元素必须和原切片的元素同类型。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了
	fmt.Println(sl3)
}

/**
 * append 在大多数情况下很好用，但是如果你想完全掌控整个追加过程，你可以实现一个这样的 AppendByte 方法：
 */
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

/**
 Go 语言中的字符串是不可变的，也就是说 str[index] 这样的表达式是不可以被放在等号左侧的。如果尝试运行 str[i] = 'D' 会得到错误：cannot assign to str[i]。

因此，您必须先将字符串转换成字节数组，然后再通过修改数组中的元素值来达到修改字符串的目的，最后将字节数组转换回字符串格式。

例如，将字符串 "hello" 转换为 "cello"：

s := "hello"
c := []byte(s)
c[0] = 'c'
s2 := string(c) // s2 == "cello"
所以，您可以通过操作切片来完成对字符串的操作。
*/
