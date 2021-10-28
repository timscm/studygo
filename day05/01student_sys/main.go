package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("welcome sms!")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	4. 删除学生
	5. 退出
	`)
}

func main() {

	var mgr manager
	mgr.allStudents = make(map[int64]*student, 40)

	for {
		showMenu()
		fmt.Print("请输入操作序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你输入的是:%d\n", choice)
		switch choice {
		case 1:
			mgr.showStudents()
		case 2:
			mgr.addStudent()
		case 4:
			mgr.deleteStudent()
		case 5:
			fmt.Println("退出")
			os.Exit(1)
		default:
			fmt.Println("输入错误")
		}
	}
}
