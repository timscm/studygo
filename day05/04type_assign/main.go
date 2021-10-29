package main

import "fmt"

// 类型断言： 具体类型
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串:", str)
	}
}

// 类型断言：通用类型
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int: ", t)
	case bool:
		fmt.Println("是一个bool: ", t)
	default:
		fmt.Println("默认类型，都没有匹配到", t)
	}
}

func main() {
	assign(100)
	assign("me...")

	assign2(100)
	assign2("me...")
	assign2(false)
	assign2([...]int{1, 2, 3})
}
