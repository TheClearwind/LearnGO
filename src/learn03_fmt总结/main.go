package main

import "fmt"

func main() {
	//a := 100
	////%T 查看类型
	//fmt.Printf("%T\n", a)
	////%v 查看数值
	//fmt.Printf("%v\n", a)
	////%b 查看二进制
	//fmt.Printf("%b\n", a)
	////%d 查看十进制
	//fmt.Printf("%d\n", a)
	////%o 查看八进制
	//fmt.Printf("%o\n", a)
	////%x 查看十六进制
	//fmt.Printf("%x\n", a)
	////%s 字符串
	//fmt.Printf("%s\n", "hello go")

	//获取用户输入

	//var s string
	//fmt.Scan(&s)
	//fmt.Println("用户输入的是:",s)
	var a, b, c int
	fmt.Scanf("%d %d %d\n", &a, &b, &c) //用户需要按照1 2 3 回车这样格式输入
	fmt.Println(a, b, c)

}
