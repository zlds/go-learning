package error

import (
	"errors"
	"testing"
)

// 不能是负数
var ErrNegative = errors.New("不能为负数")

// 超出范围
var ErrOutOfRange = errors.New("超出范围")

// Go语言中是通过if判断来处理错误的，并不像Java那样通过try-catch来处理异常。if的效率要高于try-catch。
func GetNumber(n int) (int, error) {
	if n < 0 {
		return n, ErrNegative
	}
	if n > 100 {
		return n, ErrOutOfRange
	}
	return n, nil
}

func TestError(t *testing.T) {
	if v, err := GetNumber(-1); err != nil {
		if errors.Is(err, ErrNegative) {
			// 逻辑处理
			t.Log("不能为负数。。。。。。。")
		}
	} else {
		t.Log(v, err)
	}
	if v, err := GetNumber(101); err != nil {
		if errors.Is(err, ErrOutOfRange) {
			// 逻辑处理
			t.Log("超出范围。。。。。。。")
		}
	} else {
		t.Log(v, err)
	}

	if v, err := GetNumber(99); err == nil {
		t.Log(v, err)
	}
}
