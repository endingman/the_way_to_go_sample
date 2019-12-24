package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"
	// 将整个文件的内容读到一个字符串里:io/ioutil 包里的 ioutil.ReadFile() 方法，该方法第一个返回值的类型是 []byte，里面存放读取到的内容，第二个返回值是错误，如果没有错误发生，第二个返回值为 nil
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	// 函数 WriteFile() 可以将 []byte 的值写入文件
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}

/**
 * 在很多情况下，文件的内容是不按行划分的，或者干脆就是一个二进制文件。在这种情况下，ReadString() 就无法使用了，我们可以使用 bufio.Reader 的 Read()，它只接收一个参数：

buf := make([]byte, 1024)
...
n, err := inputReader.Read(buf)
if (n == 0) { break}

变量 n 的值表示读取到的字节数.
*/
