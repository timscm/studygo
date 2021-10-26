package main

import "fmt"

func yuanfen(name string) {
	fmt.Println("Hello", name)
}

// argument is func
func lixiang(f func(string), name string) {
	fmt.Println("enter lixiang")
	defer func() {
		fmt.Println("leave lixiang")
	}()
	f(name)
}

// return func
func zhulin() func(int, int) int {
	return func(x, y int) (ret int) {
		ret = x + y
		return
	}
}

func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func low(f func()) {
	f()
}

func main() {
	lixiang(yuanfen, "元帅")
	f := zhulin()
	fmt.Println(f(1, 2))

	low(bi(yuanfen, "shuaiqi"))
}
