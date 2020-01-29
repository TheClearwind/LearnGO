package main

import "fmt"

//结构体 匿名字段
type Test struct {
	int
	string //很无聊的特性
}
type Address struct {
	province string
	city     string
}

// 嵌套结构体
type Person struct {
	name    string
	age     int
	gender  int
	Address //会自动创建一个address字段
}

func main() {
	p := Person{name: "王兰花", age: 18, gender: 1, Address: Address{"山东", "济南"}}
	fmt.Println(p.province) //可以直接访问到province
}
