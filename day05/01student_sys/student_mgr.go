package main

import "fmt"

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

type manager struct {
	allStudents map[int64]*student
}

func (m *manager) showStudents() {
	for k, v := range m.allStudents {
		fmt.Printf("学号：%d，姓名：%s\n", k, v.name)
	}
}

func (m *manager) addStudent() {
	var (
		id   int64
		name string
	)

	fmt.Print("输入学号：")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Printf("输入不是有效值:%v\n", err)
		return
	}

	fmt.Print("输入姓名：")
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Printf("输入不是有效值：%v\n", err)
		return
	}
	if _, ok := m.allStudents[id]; ok {
		fmt.Println("学号已经存在")
		return
	}
	m.allStudents[id] = newStudent(id, name)
}

func (m *manager) deleteStudent() {
	fmt.Print("输入待删除的学号：")
	var id int64
	fmt.Scanln(&id)
	if _, ok := m.allStudents[id]; !ok {
		fmt.Println("学号不存在")
		return
	}
	delete(m.allStudents, id)
}
