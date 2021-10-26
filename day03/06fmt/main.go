package main

import "fmt"

func main() {
	// 	tim@timscm:06fmt$ go run main.go
	// heh hehe
	// 用户输入的内容是: heh
	// tim@timscm:06fmt$ hehe
	// hehe: command not found
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是:", s)

	// tim@timscm:06fmt$ go run main.go
	// znz 10 22  geg
	// znz 10 22
	// tim@timscm:06fmt$ eg
	// Command 'eg' not found, but can be installed with:
	// sudo apt install easygit
	var (
		name  string
		age   int
		class string
	)
	// fmt.Scanf("%s %d %s\n", &name, &age, &class)
	// fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
