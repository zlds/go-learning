package _func

import (
	"fmt"
	"math/rand"
	"testing"
)

// 多返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

// 返回值和错误结合
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("参数不能为0")
	}
	return a / b, nil
}

func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	// 自动返回
	return
}

func TestFn1(t *testing.T) {
	i, err := divide(1, 0)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(i)
	}
	sum, product := calculate(2, 3)
	t.Log(sum, product)
}
