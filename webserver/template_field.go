package main

import (
	"fmt"
	"os"

	"text/template"
)

type Person struct {
	Name                string
	nonExportedAgeField string // 译者添加： 原文中没有定义这个，代码会报错
}

func main() {

	t := template.New("hello")

	t, _ = t.Parse("hello {{.Name}}!")

	p := Person{Name: "Mary", nonExportedAgeField: "31"} // data

	if err := t.Execute(os.Stdout, p); err != nil {

		fmt.Println("There was an error:", err.Error())

	}

}
