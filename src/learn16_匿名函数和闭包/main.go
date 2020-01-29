package main

import "fmt"

// 闭包
func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	var ret = adder(100)
	fmt.Println(ret(200))

	fmt.Println("****************************")
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9 执行完f1后 包内base 变成了11 所以再执行f2时结果为9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
