package main

import "fmt"

// 函数
/* func name(params) returntype{
}*/

func sum(a int, b int) int {
	return a + b
}

// 一种很怪的写法

func foo() (ret int) {
	return //意思就是return ret
}

//go 语言支持多返回值 很骚
func foo2() (int, string) {
	return 1, "hello"
}

// 变长参数 没有默认参数这个概念 变长参数必须放在后面
func foo3(a ...int) {
	fmt.Println(a)
}
func main() {
	foo3(1, 2, 3, 4, 5, 6)
}
