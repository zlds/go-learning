package lsp

import "testing"

type Speaker interface {
	Speak() string
}

type Animal struct {
	Name string
}

// Speak Animal实现Speaker接口
func (a Animal) Speak() string {
	return "动物叫"
}

type Dog struct {
	Animal
}

// Speak Dog实现Speaker接口
func (d Dog) Speak() string {
	return "汪汪汪"
}
func TestAnimal(t *testing.T) {
	var s Speaker
	// 使用Animal替换Speaker
	s = Animal{"其它动物"}
	t.Log(s.Speak())
	// 使用Dog替换Speaker
	s = Dog{Animal{Name: "柯基"}}
	t.Log(s.Speak())
}
