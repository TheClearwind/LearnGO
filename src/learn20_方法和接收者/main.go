package main

import "fmt"

//Person 是一个结构体
type Person struct {
	name, city string
	age        uint8
}

//构造函数
func newPerson(name, city string, age uint8) *Person {
	return &Person{name: name,
		city: city,
		age:  age, //如果这样换行写 就得最后加个逗号
	}
}

//定义方法
func (p Person) Say(content string) { //p Person 为接收者 也就是说是Person的方法 类似于python的self
	fmt.Printf("%s say %s\n", p.name, content)
}

func (p *Person) SetAge(age uint8) {
	p.age = age // 当需要修改值的时候 要使用指针接收者
}
func main() {
	p := newPerson("王兰花", "北京", 18)
	p.Say("诸葛大力")

	fmt.Println(p.age)
	p.SetAge(17)
	fmt.Println(p.age)
}
