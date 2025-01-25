package nested

import (
	"testing"
)

type A interface {
	test1()
}

type B interface {
	test2()
}

// 嵌套接口，接口 C 包含接口 A 和 B 的所有方法。如果要实现C接口必须实现A和B接口的所有方法
type C interface {
	A
	B
	test3()
}

type Person struct {
}

func (p Person) test1() {
	println("test1")
}

func (p Person) test2() {
	println("test2")
}

func (p Person) test3() {
	println("test3")
}

func TestInterface(t *testing.T) {
	var person Person = Person{}
	// 实现 C 接口的所有方法
	person.test1()
	person.test2()
	person.test3()
}
