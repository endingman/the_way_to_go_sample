package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

/*这个结构会保存解析后的返回数据。
他们会形成有层级的XML，可以忽略一些无用的数据*/
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func main() {
	// 发起请求查询推特Goodland用户的状态
	response, _ := http.Get("http://twitter.com/users/Googland.xml")
	// 初始化XML返回值的结构
	user := User{xml.Name{"", "user"}, Status{""}}
	// 将XML解析为我们的结构
	// 你可能无法获取到 xml 数据，另外由于 go 版本的更新，xml.Unmarshal 函数的第一个参数需是 [] byte 类型，而无法传入 Body。
	xml.Unmarshal(response.Body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
