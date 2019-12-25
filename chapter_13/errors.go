// errors.go
package main

import (
	"errors"
	"fmt"
)

//任何时候当你需要一个新的错误类型，都可以用 errors（必须先 import）包的 errors.New 函数接收合适的错误信息来创建，像下面这样：
var errNotFound error = errors.New("Not found error")

/**
 Go 有一个预先定义的 error 接口类型

type error interface {
    Error() string
}
 * @Author Liumm
 * @Date   2019-09-04
 * @return {[type]}   [description]
*/
func main() {
	fmt.Printf("error: %v", errNotFound)
}

// error: Not found error
