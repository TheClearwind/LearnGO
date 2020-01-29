package main

import "fmt"

// 类型别名
type myint = int

// 自定义类型
type myint2 int //没有等号了

// 结构体
type Persion struct {
	name   string
	age    int
	gender int
	hobby  []string
	eat    func()
	say    func(string)
}

//伪造构造函数
func newPerson(name string, age, gender int) *Persion { //返回指针是为了节省复制的开销 因为结构体是值类型
	return &Persion{name: name, age: age, gender: gender}
}
func main() {
	//var a myint
	//a = 5
	//fmt.Printf("%T\n", a) //int
	//var b myint2
	//b = 5
	//fmt.Printf("%T\n", b) //main.myint2
	//
	//var p Persion
	//p.name = "王兰花"
	//p.age = 18
	//p.gender = 0
	//p.hobby = []string{"你懂的", "钱"}
	//p.eat = func() {
	//	fmt.Println("吃饭ing")
	//}
	//p.say = func(s string) {
	//	fmt.Println("say:", s)
	//}
	//fmt.Println(p)
	//p.eat()
	//fmt.Printf("%T\n",p)

	//匿名结构体
	var user struct{ name, pwd string }
	user.name = "王兰花"
	user.pwd = "123456"

	//结构体指针
	//p := new(Persion) //p为 Persion指针 但是可以直接点出字段和方法
	//fmt.Printf("%T\n", p)
	//p.age = 18
	//p.name = "王兰花"
	//fmt.Printf("%v", p.age)

	// 结构体初始化
	p3 := Persion{name: "王兰花", age: 18} //冒号 赋值
	fmt.Printf("%#v", p3)
}
