package main

/*

export -n GOPROXY
unset GOPROXY

# 方法一：依赖的库已经在github了
go get github.com/timscm/studygo/calc
go mod init xx
go mod tidy
go run main.go
# 生成了 go.mod, go.sum

# 方法二：依赖的库还在本地
go mod init xx
go mod edit -replace github.com/timsc/studygo/calc=../../calc
go mod tidy
go run main.go
# 只生成 go.mod

*/

import (
	"fmt"

	"github.com/timscm/studygo/calc"
)

func main() {
	ret := calc.Add(1, 2)
	fmt.Println("ret: ", ret)
}
