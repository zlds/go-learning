package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到取消信号，退出")
			return
		default:
			fmt.Println(name, "工作中.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "work1")
	go work(ctx, "work2")
	time.Sleep(2 * time.Second)
	fmt.Println("通知goroutine退出")
	// 触发ctx.Done()，通知所有goroutine退出
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("主程序退出")
}
