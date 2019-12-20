package main

import "fmt"

/**在函数块里面，return 之后的语句都不会执行。如果一个函数需要返回值，那么这个函数里面的每一个代码分支（code-path）都要有 return 语句。**/

// func (st *Stack) Pop() int {
//     v := 0
//     for ix := len(st) - 1; ix >= 0; ix-- {
//         if v = st[ix]; v != 0 {
//             st[ix] = 0
//             v = st[ix]
//             return v  //==>在函数块里面，return 之后的语句都不会执行,v不会执行
//         }
//     }

//     return v
// }

func main() {
    fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
    // var i1 int = MultiPly3Nums(2, 5, 6)
    // fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
}

func MultiPly3Nums(a int, b int, c int) int {
    // var product int = a * b * c
    // return product
    return a * b * c
}
