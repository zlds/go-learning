package base

import (
	"fmt"
	"testing"
)

// 空接口可以代表任何类型，相当于Java中的Object。在使用的时候并不需要强制类型转换，Go语言通过断言来实现类型转换。
func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer: ", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("String: ", i)
		return
	}
	fmt.Println("Unknown Type")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("hello")
}
