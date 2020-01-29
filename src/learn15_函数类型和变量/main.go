package main

import "fmt"

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 1
}

// 函数也可以作为参数传递
func f3(x func() int) {
	fmt.Println(x())
}

//函数也可以作为返回值
func f4() func(int, int) int {
	return func(x, y int) (ret int) { return x + y }
}
func main() {
	// 函数也是一种类型
	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n", b)

	f3(b)
	fmt.Println(f4()(1, 2))
}
