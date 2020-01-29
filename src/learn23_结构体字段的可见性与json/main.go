package main

import (
	"encoding/json"
	"fmt"
)

//在go 语言中 定义的标识符 首字母大写则对外可见

type student struct {
	Id   string
	Name string
}

func newStudent(id, name string) *student {
	return &student{Id: id, Name: name}
}

//这里好坑人啊 字段的首字母如果不大写 被认为无法导出，所以不会参与json序列化 坑人！
type class struct {
	//struct 可通过tag 为某个包指明 字段标识符
	Id       string `json:"class_name"` //json包中会将Id 识别为 class_name
	Students []student
}

func main() {
	c1 := class{Id: "高三五班", Students: make([]student, 0, 30)}
	for i := 0; i < 10; i++ {
		c1.Students = append(c1.Students, *newStudent(fmt.Sprintf("%02d", i), fmt.Sprintf("stu_%02d", i)))
	}
	//data, err := json.Marshal(c1)
	//if err == nil {
	//	fmt.Printf("%s", data)
	//} else {
	//	fmt.Println("json序列化失败！")
	//}
	//
	//反序列化
	s := `{"class_name":"高三五班","Students":[{"Id":"00","Name":"王兰花"},{"Id":"01","Name":"stu_01"}]}`
	var temp class
	err := json.Unmarshal([]byte(s), &temp)
	if err != nil {
		fmt.Println("反序列化失败", err)
	} else {
		fmt.Printf("%v", temp)
	}

}
