package main

import (
	"fmt"
)

// 全局作用域
var x = 100 // 定义一个全局变量

func f1() {
	// x := 10 // 局部作用域
	fmt.Println(x)
}

func main() {
	f1() // 100

	// 语句块作用域
	if i := 10; i < 18 {
		fmt.Println("good good study")
	}
	// fmt.Println(i) 无法访问 i

	// 语句块作用域
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	// fmt.Println(j)

}
