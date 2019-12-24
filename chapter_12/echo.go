package main

import (
	/**
	flag 包有一个扩展功能用来解析命令行选项。但是通常被用来替换基本常量，例如，在某些情况下我们希望在命令行给常量一些不一样的值。（参看 19 章的项目）

	在 flag 包中一个 Flag 被定义成一个含有如下字段的结构体：

	type Flag struct {
	    Name     string // name as it appears on command line
	    Usage    string // help message
	    Value    Value  // value as set
	    DefValue string // default value (as text); for usage message
	}
	*/
	"flag" // command line option parser
	"os"
)

// flag.Bool() 定义了一个默认值是 false 的 flag：当在命令行出现了第一个参数（这里是 "n"），flag 被设置成 true（NewLine 是 *bool 类型）。flag 被解引用到 *NewLine，所以当值是 true 时将添加一个 newline（"\n"）
var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	// flag.Parse() 扫描参数列表（或者常量列表）并设置 flag, flag.Arg(i) 表示第 i 个参数
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	//flag.Narg() 返回参数的数量。解析后 flag 或常量就可用了。
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		// Parse() 之后 flag.Arg(i) 全部可用，flag.Arg(0) 就是第一个真实的 flag，而不是像 os.Args(0) 放置程序的名字
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
