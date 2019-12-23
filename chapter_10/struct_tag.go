package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	// 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记。
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	// 标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它
	ttType := reflect.TypeOf(tt)

	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
