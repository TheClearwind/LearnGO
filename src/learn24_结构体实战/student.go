package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}
type studentManager struct {
	allstus []*student
}

func newManager(c int) *studentManager {
	return &studentManager{allstus: make([]*student, 0, c)}
}

//构造函数
func newStudent(id int, name, class string) *student {
	return &student{id: id, name: name, class: class}
}

func (sm *studentManager) addStudent(id int, name, class string) {
	sm.allstus = append(sm.allstus, newStudent(id, name, class))
	fmt.Println("添加成功！")
}

func (sm *studentManager) modifyStudent(stu *student) {
	for i, v := range sm.allstus {
		if v.id == stu.id {
			sm.allstus[i] = stu
			fmt.Println("修改成功")
			return
		}
	}
	fmt.Printf("学号：%05d 查无此人！", stu.id)
}
func (sm *studentManager) showStudents() {
	fmt.Println(len(sm.allstus))
	for _, v := range sm.allstus {
		fmt.Printf("班级:%s 学号:%05d 姓名:%s\n", v.class, v.id, v.name)
	}
}
