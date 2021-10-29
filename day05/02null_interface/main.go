package main

import "fmt"

// 空接口

// interface: 关键字
// interface{}: 空接口类型 === type x interface{} === 匿名声明类型
func useNULLInterface() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zhoulin"
	m1["age"] = 9000
	m1["abc"] = true
	m1["hobby"] = [...]string{"hao", "yinyue", "rap"}
	fmt.Println(m1)
}

// 空接口作为参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	useNULLInterface()
	show(false)
	show(nil)
	show(10)
	show("me")
}
