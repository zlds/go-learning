package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func task(id int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Task %d running: %d\n", id, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func TestGoroutine(t *testing.T) {
	// 设置P的数量(CPU的核数)
	runtime.GOMAXPROCS(2)
	for i := 1; i <= 5; i++ {
		// 创建5个协程
		go task(i)
	}
	time.Sleep(time.Second * 2)

}
