package fib

import "testing"

func TestFib(t *testing.T) {

	var a int = 1
	var b int = 1

	for i := 0; i < 5; i++ {
		t.Log(a)
		tmp := a
		a = b
		b = tmp + a
	}
}
