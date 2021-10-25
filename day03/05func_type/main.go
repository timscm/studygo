package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("Hello world")
}

// x是函数类型, 函数也可以作为参数进行传递
// x = func() int {} 匿名函数
func f2(x func() int) int {
	return 1
}

// 函数作为返回值
func f3(x func() int) func(int) int {
	return func(int) int {
		fmt.Println("return func")
		return 33
	}
}

func main() {
	a := f1
	fmt.Printf("%T\n", a) // func()
	b := f2
	fmt.Printf("%T\n", b) // func(func() int) int

	// c 是一个函数类型的变量
	var c func(int) int

	// 给函数类型的变量赋值
	c = f3(func() int {
		return 3
	})
	fmt.Println(c(1))
}
