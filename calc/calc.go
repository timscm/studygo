package calc

/*
	main (import A)
	A (import B)
	B (import C)

	init 顺序：
	C.init()
	B.init()
	A.init()
	main.init()

*/

import "fmt"

var x int = 10

func init() {
	fmt.Println("calc init 1...", x, y)
}

func init() {
	fmt.Println("calc init 2...", x, y)
}

func init() {
	fmt.Println("calc init 3...", x, y)
}

func Add(x, y int) int {
	return x + y
}

var y int = 20
