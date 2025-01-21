package _map

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T) {
	// map的value可以是一个方法。value的类型为func(op int) int
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	t.Log(m[1](2), m[2](2))
}
