package main

import (
	"fmt"
	"go-learning/constant"
	"go-learning/example"
)

func main() {
	fmt.Println("Hello Word!")
	example.SayHello()

	// 打印常量
	fmt.Println(constant.Count)
	// 打印iota
}
