// gob2.go
package main

import (
	"encoding/gob"
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

var content string

/**
 * Gob 使用通用的 io.Writer 接口，通过 NewEncoder() 函数创建 Encoder 对象并调用 Encode()；相反的过程使用通用的 io.Reader 接口，通过 NewDecoder() 函数创建 Decoder 对象并调用 Decode
 * @Author Liumm
 * @Date   2019-09-03
 * @return {[type]}   [description]
 */
func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
