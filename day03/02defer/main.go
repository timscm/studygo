package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	// defer 按照声明顺序，先进后出：栈
	defer fmt.Println("in the end 1")
	defer fmt.Println("in the end 2")
	defer fmt.Println("in the end 3")
	fmt.Println("not real end.")
}

func main() {
	deferDemo()
}
