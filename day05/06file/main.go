package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func osopenread() {
	// Open 不会创建文件
	fd, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer fd.Close()

	var tmp [128]byte
	for {
		n, err := fd.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("read finished(io.EOF)")
			return
		}

		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(tmp[:n]))
		// if n < 128 {
		// 	fmt.Println("read finished.(size)")
		// 	return
		// }
	}
}

func bufioread() {
	fd, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file faile,d err:%v\n", err)
		return
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read file end.(io.EOF)")
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	// osopenread()
	bufioread()
}
