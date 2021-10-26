package main

import "fmt"

// 递归函数：自己调用自己
// 递归适合处理那种问题相同，问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 永远不要高估自己

// 上台阶的问题
// n个台阶，一次可以走1步，也可以走2布，有多少种走法
func f2(n int) int {
	if n <= 2 {
		return n
	}
	return f2(n-1) + f2(n-2)
}

// 阶乘
// 3! = 3*2*1     = 3*2!
// 4! = 4*3*2*1   = 4*3!
// 5! = 5*4*3*2*1 = 5*4!
func f(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	ret := f(6)
	fmt.Println(ret)

	fmt.Println(f2(5))
}
