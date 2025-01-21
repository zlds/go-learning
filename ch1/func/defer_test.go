package _func

import (
	"testing"
)

// defer 可用于资源清理(释放资源，释放锁)
func TestDefer(t *testing.T) {
	// 创建匿名函数
	defer func() {
		t.Log("Clear resources")
	}()
	t.Log("Started")
	// defer 仍会执行
	panic("Fatal error")
}
