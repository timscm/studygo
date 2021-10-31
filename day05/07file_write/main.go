package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容

func oswritefile() {
	// _, err := os.OpenFile("./xx.txt", os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Printf("open file failed, err:%v\n", err)
	// 	return
	// }

	// fd, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	fd, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fd.Close()

	n, err := fd.Write([]byte("123\n"))
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	fmt.Println(n)

	n, err = fd.WriteString("new\n")
	if err != nil {
		fmt.Printf("writeString file failed, err:%v\n", err)
		return
	}
	fmt.Println(n)
}

func bufiowritefile() {
	// bufio.NewWriter
	fd, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer fd.Close()

	wr := bufio.NewWriter(fd)
	wr.WriteString("hello, good boy\n") // 写到缓存中
	wr.Flush()                          // 将缓存中的内容写入文件
}

func ioutilwritefile() {
	content := "hello ioutil\n"
	// 覆盖了文件的内容
	err := ioutil.WriteFile("./xx.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("write file failed", err)
		return
	}
}

func main() {
	oswritefile()
	bufiowritefile()
	ioutilwritefile()
}
