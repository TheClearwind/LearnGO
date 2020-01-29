package main

import "fmt"

func main() {

	//Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和*（根据地址取值）\

	//var a = 50
	//b:=&a
	//fmt.Println(*b)
	//*b=80
	//fmt.Println(a)

	// make 用来给引用类型 申请内存
	//new  用来给值类型申请内存
	var a []int // 只声明是不能使用的
	//使用make 初始化
	a = make([]int, 3, 3)
	fmt.Println(a)
}
