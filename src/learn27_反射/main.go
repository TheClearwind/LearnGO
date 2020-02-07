package main

import (
	"fmt"
	"reflect"
)

//通过反射获取数据类型

func getType(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
}

//通过反射获取值
func reflectValue(arg interface{}) {
	v := reflect.ValueOf(arg)
	fmt.Printf("%v %T\n", v, v)
}

// 通过反射设置值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind() //x 为一个指针 所以需要使用Elem
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	}
}

type dog struct {
}
type myint int

//  结构体反射

type student struct {
	Name  string `json:"name"` // 首字母要大写否则反射无法拿到
	Score int
}

func (s student) Say(content string) { // 首字母要大写否则反射无法拿到
	fmt.Printf("%s say %s\n", s.Name, content)
}
func main() {
	//var a int32=30
	//reflectSetValue(&a)
	//fmt.Println(a)

	stu := student{Name: "王兰花", Score: 80}
	//t := reflect.TypeOf(stu)
	//for i:=0;i<t.NumField();i++{
	//	field := t.Field(i)
	//	fmt.Printf("Name:%v Type:%v Tag:%v\n",field.Name,field.Type,field.Tag)
	//}

	//通过名字取字段
	//field, ok := t.FieldByName("Name")
	//if ok {
	//	fmt.Printf("Name:%v Type:%v Tag:%v\n", field.Name, field.Type, field.Tag)
	//} else {
	//	fmt.Println("不存在")
	//}

	// 获取方法并且调用
	v := reflect.ValueOf(stu)
	fmt.Println(v.NumMethod())
	method := v.MethodByName("Say")
	arg := []reflect.Value{reflect.ValueOf("hello world")}
	method.Call(arg)
	fmt.Println(method.Type())
}
