package error

import (
	"errors"
	"testing"
	"time"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func TestErrorPerformance(t *testing.T) {
	count := 1000000
	// 正常测试
	start := time.Now()
	for i := 0; i < count; i++ {
		_, err := divide(10, 2)
		if err != nil {

		}
	}
	t.Log("正常测试耗时：", time.Since(start))

	// 异常测试
	start = time.Now()
	for i := 0; i < count; i++ {
		_, err := divide(10, 0)
		if err != nil {

		}
	}
	t.Log("异常测试耗时：", time.Since(start))
}
