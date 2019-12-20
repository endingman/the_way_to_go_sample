package main

import (
	"flag" // command line option parser
	"os"
)

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
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
