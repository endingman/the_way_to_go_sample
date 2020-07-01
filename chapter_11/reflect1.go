// blog: Laws of Reflection
package main

import (
	"fmt"
	"reflect"
)

/**
 两个简单的函数，reflect.TypeOf 和 reflect.ValueOf，返回被检查对象的类型和值。例如，x 被定义为：var x float64 = 3.4，那么 reflect.TypeOf(x) 返回 float64，reflect.ValueOf(x) 返回 <float64 Value>

实际上，反射是通过检查一个接口的值，变量首先被转换成空接口。
 * @Author Liumm
 * @Date   2019-12-23
 * @return {[type]}   [description]
*/
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) //reflect.TypeOf返回被检查对象的类型
	v := reflect.ValueOf(x)                 //reflect.ValueOf返回被检查对象的值
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind()) //Kind 总是返回底层类型
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
