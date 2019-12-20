package main

import (
	"bufio"
	"compress/gzip" //提供压缩功能的包
	"fmt"
	"os"
)

func main() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	// 打开文件
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	// bufio.NewReader 来获得一个读取器变量。

	// 通过使用 bufio 包提供的读取器
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		// ReadString('\n') 或 ReadBytes('\n') 将文件的内容逐行（行结束符 '\n'）读取出来
		// Unix 和 Linux 的行结束符是 \n，而 Windows 的行结束符是 \r\n。在使用 ReadString 和 ReadBytes 方法的时候，我们不需要关心操作系统的类型，直接使用 \n 就可以了。另外，我们也可以使用 ReadLine() 方法来实现相同的功能。
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
