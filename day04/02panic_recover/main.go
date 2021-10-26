package main

import "fmt"

func f1() {
	defer func() {
		err := recover()
		fmt.Println("in defer ...")
		fmt.Println(err)
		fmt.Println("leave defer ...")
	}()
	panic("犯了不可原谅的错误")
	// fmt.Println("f1") // panic 之后的代码不会执行
}

func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()
}
