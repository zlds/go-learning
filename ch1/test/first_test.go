package test

import "testing"

func TestFirst(t *testing.T) {
	t.Log("My first test")
}

// 变量交换
func TestExchange(t *testing.T) {
	var a int = 1
	var b int = 2
	a, b = b, a
	t.Log(a, b)
}
