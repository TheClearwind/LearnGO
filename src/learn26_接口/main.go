package main

import "fmt"

// 接口 -> 通用类型
type dog struct {
}
type cat struct {
}

type sayer interface {
	say() // 实现了say方法的 都叫sayer 类型  实现全部的方法 才行！！！
	//定义 方法 方法名（参数列表) 返回值
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵")
}

func do(sayer sayer) {
	sayer.say()
}

// 空接口  任何类型 都是空接口类型 空接口 不需要提前定义
var x interface{} // 定义一个空接口变量 可以接受任意值

//空接口应用
//1. 作为函数参数
//2. map的值类型

func main() {
	d := dog{}
	c := cat{}
	do(d)
	do(c)

	//类型断言
	ret, ok := x.(int)
	if !ok {
		fmt.Println("不是int")
	} else {
		fmt.Println(ret)
	}

	// type switch 检测类型
	switch v := x.(type) {
	case int:
		fmt.Println("int 类型 value:", v)
	default:
		fmt.Println(v)
	}
}
