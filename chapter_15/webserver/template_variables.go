package main

import (
	"os"

	"text/template"
)

func main() {

	t := template.New("test")
	// 你可以在变量名前加一个「$」符号来为模板中的管道创建一个局部变量。变量名称只能由字母、数字、下划线组成。
	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))

	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))

	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}} {{end}}!\n"))

	t.Execute(os.Stdout, nil)

}
