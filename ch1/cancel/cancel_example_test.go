package cancel

import (
	"fmt"
	"testing"
	"time"
)

var stop bool

// 通过共享变量的方式通知goroutine退出(多个goroutine共享变量需要加锁，如果goroutine在sleep或阻塞，它不会立刻退出)
func worker() {
	for {
		//
		if stop {
			fmt.Println("worker 收到退出信号")
			return
		}
		fmt.Println("worker 运行中.")
		time.Sleep(500 * time.Millisecond)
	}
}

func workerChannel(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("workerChannel 收到退出信号")
			return
		default:
			fmt.Println("workerChannel 运行中.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func TestStop(t *testing.T) {
	// 共享变量方式实现goroutine退出
	go worker()
	time.Sleep(2 * time.Second)
	stop = true
	time.Sleep(1 * time.Second)

	// 通过channel方式实现goroutine退出(不支持子goroutine继承上下文)
	stopCh := make(chan struct{})
	go workerChannel(stopCh)
	time.Sleep(2 * time.Second)
	close(stopCh)
	time.Sleep(1 * time.Second)
	fmt.Println("主程序退出")

	// 最佳实践：使用context实现goroutine退出(参考cancel_test.go示例)

}
