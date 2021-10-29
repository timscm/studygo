package main

import "fmt"

// 接口：
// +----------------+
// |  动态类型       |
// +----------------+
// |   动态值        |
// +----------------+
//  值接收者：动态类型可以是值类型，指针类型
// 指针接收者：动态类型值可以是指针类型

// 值接收者实现接口的方法
// vs
// 指针接收者实现接口的方法

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 值接收者实现了接口animal的方法move
// var a1 animal
// a1 = cat{ // 也可以&cat
//   name: "tom",
//   feet: 4,
// }
// func (c cat) move() {
// 	fmt.Println("走猫步...")
// }

// func (c cat) eat(food string) {
// 	fmt.Printf("猫吃%s...\n", food)
// }

// 指针接收者实现了接口animal的方法move
// var a1 animal
// a1 = &cat{ // 只能是cat
//   name: "tom",
//   feet: 4,
// }
func (c *cat) move() {
	fmt.Println("走猫步...")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 animal

	c1 := cat{"tom", 4}
	c2 := cat{"假老练", 4}

	// 值接收者实现的接口animal方法
	// a1 = c1
	// fmt.Println(a1)
	// a1 = &c2
	// fmt.Println(a1)

	// 指针接收者实现的接口animal方法
	a1 = &c1
	fmt.Println(a1)
	a1 = &c2
	fmt.Println(a1)
}
