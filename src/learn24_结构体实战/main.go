package main

import (
	"fmt"
	"os"
)

//go 语言支持多返回值
func getInput() (int, string, string) {
	var (
		id    int
		name  string
		class string
	)
	fmt.Print("请输入学号:")
	fmt.Scanf("%d\n", &id)

	fmt.Print("请输入姓名:")
	fmt.Scanf("%s\n", &name)

	fmt.Print("请输入班级:")
	fmt.Scanf("%s\n", &class)
	return id, name, class
}

//学员信息管理系统
//增删改查
func main() {
	manager := newManager(30)
	for {
		showMenu()

		var input = 0
		fmt.Print("您的选择是:")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			id, name, class := getInput()
			manager.addStudent(id, name, class)
		case 2:
			id, name, class := getInput()
			manager.modifyStudent(newStudent(id, name, class))
		case 3:
			manager.showStudents()
		case 4:
			os.Exit(0)
		}
	}

}

func showMenu() {
	fmt.Println("欢迎使用正方管理系统：\n1.添加学员\n2.编辑学员信息\n3.显示学员信息\n4.退出")
}
