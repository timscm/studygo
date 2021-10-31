package main

import (
	"fmt"
	"os"
)

func main() {
	// Open 不会创建文件
	fd, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer fd.Close()

	var tmp [1024]byte
	n, err := fd.Read(tmp[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		return
	}
	fmt.Printf("read %d bytes\n", n)
	fmt.Println(string(tmp[:n]))
}
