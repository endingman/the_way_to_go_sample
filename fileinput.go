package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
 * 在 Go 语言中，文件使用指向 os.File 类型的指针来表示的，也叫做文件句柄。我们在前面章节使用到过标准输入 os.Stdin 和标准输出 os.Stdout，他们的类型都是 *os.File。
 * @Author Liumm
 * @Date   2019-08-26
 * @return {[type]}   [description]
 */
func main() {
	// 变量 inputFile 是 *os.File 类型的。该类型是一个结构，表示一个打开文件的描述符（文件句柄）。然后，使用 os 包里的 Open 函数来打开一个文件。该函数的参数是文件名，类型为 string。
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		// Unix 和 Linux 的行结束符是 \n，而 Windows 的行结束符是 \r\n
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
