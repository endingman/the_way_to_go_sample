// json.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	// json.Marshal() 的函数签名是 func Marshal(v interface{}) ([]byte, error)

	/**
	   出于安全考虑，在 web 应用中最好使用 json.MarshalforHTML() 函数，其对数据执行 HTML 转码，所以文本可以被安全地嵌在 HTML <script> 标签中。

	  json.NewEncoder() 的函数签名是 func NewEncoder(w io.Writer) *Encoder，返回的 Encoder 类型的指针可调用方法 Encode(v interface{})，将数据对象 v 的 json 编码写入 io.Writer w 中。

	  JSON 与 Go 类型对应如下：

	  bool 对应 JSON 的 booleans
	  float64 对应 JSON 的 numbers
	  string 对应 JSON 的 strings
	  nil 对应 JSON 的 null
	  不是所有的数据都可以编码为 JSON 类型：只有验证通过的数据结构才能被编码：

	  JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map [string] T（T 是 json 包中支持的任何类型）
	  Channel，复杂类型和函数类型不能被编码
	  不支持循环数据结构；它将引起序列化进入一个无限循环
	  指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
	       * @type {[type]}
	*/

	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
