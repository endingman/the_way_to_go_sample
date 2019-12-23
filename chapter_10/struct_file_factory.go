package main

import "fmt"

type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

/**
 * 结构体类型对应的工厂方法，它返回一个指向结构体实例的指针
 */
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}

func main() {
	f := NewFile(10, "test.txt")
	fmt.Println(f.name)
}
