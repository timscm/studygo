package main

import "fmt"

// defer 时机：
// return语句在底层并不是原子操作。
// Go语言中函数的return不是原子操作，在底层时分为两步来执行
// 第一步：返回值赋值
// 第二步：真正的RET返回
// 函数中如果存在 defer，那么defer执行的时机时在第一步和第二步之间

// 正常情况：
// 		return x --> 返回值=x, RET指令
// defer情况：
// 		return x --> 返回值=x, 运行defer, RET指令

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	// 第一步： 返回值赋值（得到值：5）
	// 第二步： 执行defer函数 （修改的x，不是返回值）
	// 第三步： 执行RET指令
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()

	// 第一步：返回值赋值（x=5）
	// 第二步：执行defer函数（修改的x，x=6）
	// 第三步：执行RET x（返回的是x, x==6）
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	// 第一步：y=x, y=5
	// 第二步：x++, x==6
	// 第三步：RET y, 返回5
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	// 第一步：返回值=x, ==5
	// 第二步：执行defer，修改的是内部x，与外部x无关
	// 第三步：执行RET x，==5
	return 5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}
