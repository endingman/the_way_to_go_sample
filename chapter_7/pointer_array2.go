package main

import "fmt"

/**
 * 你可以取任意数组常量的地址来作为指向新实例的指针。
 * @Author Liumm
 * @Date   2019-08-20
 * @param  {[type]}   a *[3]int       [description]
 * @return {[type]}     [description]
 */
func fp(a *[3]int) {
    fmt.Println(a)
}

func main() {
    for i := 0; i < 3; i++ {
        fp(&[3]int{i, i * i, i * i * i})
    }
}

// 在这里数组长度同样可以写成 ... 或者直接忽略。
// var arrLazy = [...]int{5, 6, 7, 8, 22}
// 从技术上说它们其实变化成了切片。
