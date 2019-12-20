package main

import (
	"os"

	"text/template"
)

func main() {

	t := template.New("test")
	// 可以在你的代码中使用一些预定义的模板函数，例如： 和 fmt.Printf 函数类似的 printf 函数
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}} {{end}}!\n"))

	t.Execute(os.Stdout, nil)

}
