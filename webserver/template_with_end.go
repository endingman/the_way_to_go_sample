package main

import (
	"os"

	"text/template"
)

func main() {

	t := template.New("test")
	//在 Go 模板中使用 (.) ： 他的值 {{.}} 被设置为当前管道的值。
	//with 语句将点的值设置为管道的值。如果管道是空的，就会跳过 with 到 end 之前的任何内容；当嵌套使用时，点会从最近的范围取值。在下面这个程序中会说明：
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")

	t.Execute(os.Stdout, nil)

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}} {{end}}!\n")

	t.Execute(os.Stdout, nil)

}
