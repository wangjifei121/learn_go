package main

import "fmt"

//定义学生类型
type student struct {
	name string
	id   int
}

//学生实例化
func newStudent(name string, id int) *student {
	return &student{
		name: name,
		id:   id,
	}
}

//学员管理系统
type studentManager struct {
	allstudent []student
}

//学员管理系统实例化
func newStudentManager() *studentManager {
	stu := make([]student, 0, 100)
	return &studentManager{
		allstudent: stu,
	}
}

//添加学员
func (sm *studentManager) addStudent(name string, id int) {
	new_stu := newStudent(name, id)
	sm.allstudent = append(sm.allstudent, *new_stu)
}

//修改学员
func (sm *studentManager) editStudent(name string, id int) {
	for index, s := range sm.allstudent {
		if s.id == id {
			sm.allstudent[index].name = name
			fmt.Printf("修改学生成功:%d\n", index)
			return
		}
	}
	fmt.Println("当前输入的学员ID不存在，请重新输入")
}

//查看学员
func (sm *studentManager) showStudent() {
	fmt.Println("当前学员信息如下：")
	for index, s := range sm.allstudent {
		fmt.Printf("%d - 学号:%d  姓名:%v\n", index, s.id, s.name)
	}
}
