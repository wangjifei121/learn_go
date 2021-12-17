package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("----欢迎进入学院管理系统---")
	fmt.Printf("1: 添加学员信息\n")
	fmt.Printf("2: 修改学员信息\n")
	fmt.Printf("3: 查看学员信息\n")
	fmt.Printf("4: 退出学员管理系统\n")
}

//获取用户输入
func getInput() (string, int) {
	var id int
	var name string
	fmt.Println("请输入您的ID")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入您的name")
	fmt.Scanf("%v\n", &name)
	return name, id
}

func main() {
	sm := newStudentManager()
	
	for {
		showMenu()
		var input int
		fmt.Println("请选择您要操作的编号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("当前用户输入的信息是 ", input)
		switch input {
		case 1:
			name, id := getInput()
			sm.addStudent(name, id)
		case 2:
			var id int
			var name string
			fmt.Println("请输入您要修改的ID")
			fmt.Scanf("%d\n", &id)
			fmt.Println("您要将name修改为")
			fmt.Scanf("%v\n", &name)
			sm.editStudent(name, id)
		case 3:
			//show
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}

}
