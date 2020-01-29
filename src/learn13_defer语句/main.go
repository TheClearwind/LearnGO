package main

import "fmt"

//
//import "fmt"
//
//func deferDemo() {
//	fmt.Println("start")
//	defer fmt.Println("hello world") //defer 后面的语句会延时到函数即将返回时执行
//	defer fmt.Println("hello golang") //多个defer时从最后一个开始执行
//	fmt.Println("end")
//}
//
//func main() {
//	deferDemo()
//}

//defer 执行时机：
// go 语言中的return 其实是两步操作 return result => 1.ret = result 2.return ret
// 在含有derfer 的函数中 return result => 1.ret = result 2.defer 3.return ret
//defer 经典问题
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) { //1.x=5 2.defer -> x=6 3.return x x=6 所以返回6
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) { //1. y = x = 5 2. x++ 3. return y = 5  所以返回5
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) { //1.x = 5 2.局部变量x++ 3.返回 x =5 所以返回5
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) // defer calc("AA",1,?)->"A" 1 2 3->defer calc("AA",1,3)
	x = 10
	defer calc("BB", x, calc("B", x, y)) // defer calc("BB", 10,?) -> "B" 10 2 12 -> defer calc("BB", 10,12)
	y = 20

	//defer calc("BB", 10,12) -> BB 10 12 22
	//defer calc("AA",1,3)  -> AA 1 3 4
}
