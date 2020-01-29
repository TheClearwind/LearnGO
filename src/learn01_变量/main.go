package main

import "fmt"

func main() {
	//// 定义变量的方式1：var name type
	//var v string
	//v = "Hello world"
	//fmt.Println(v)

	//// 定义变量时 支持类型推断 可以省略type
	//var age = 18
	//fmt.Println("my age is",age)

	//// 声明多个变量
	//var name,des = "张三","帅哥"
	//fmt.Println(name," is ",des)

	// 声明不同类型的变量
	//var (
	//	name = "lambda"
	//	age  = 18
	//)
	//fmt.Println("My name is",name,"my age is",age)

	//简单声明
	name, age := "lambda", 18 //:= 可以用来赋值，但是左边必须有一个没有声明的变量 只能在函数体内使用
	fmt.Println("My name is", name, "my age is", age)
}
