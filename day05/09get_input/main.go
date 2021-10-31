package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时，如果有空格

func useScan() {
	var s string
	fmt.Print("请输入内容:")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是:%s\n", s)
}

func useBufio() {
	var s string
	// NewReader的参数就是一个接口类型
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("reader error", err)
		return
	}
	fmt.Printf("你输入的内容是:%s\n", s)
}

func main() {
	// useScan()
	useBufio()
}
