package main

import "fmt"

// 可变参数
func f5(title string, y ...int) int {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // sum 不需要声明
	return      // 返回命名参数，不需要指定参数
}

func main() {
	f5("title", 1, 2, 3, 4, 5, 6)
	f6(1, 2)

	// 在命名函数中，不能在声明命名函数
	// 但可以使用匿名函数
	x := func() {
		fmt.Println("in x")
	}
	x()
}
