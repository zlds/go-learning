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

// 手动检查context是否被取消
func worker1(ctx context.Context, name string) {
	for {
		if ctx.Err() != nil {
			fmt.Println(name, "收到取消信号，退出")
			return
		}
		fmt.Println(name, "工作中.")
		time.Sleep(500 * time.Millisecond)
	}
}

// 在执行完cancel之后会进行sleep(可通过sync.WaitGroup替代)，这是为了防止进程可能过早的退出导致ctx.Done()没有机会执行
func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "work1")
	go work(ctx, "work2")
	time.Sleep(2 * time.Second)
	fmt.Println("通知goroutine退出")
	// 触发ctx.Done()，通知所有goroutine退出
	cancel()
	time.Sleep(1 * time.Second)

	// 手动检查ctx.Err()，通知goroutine退出
	ctx1, cancel1 := context.WithCancel(context.Background())
	go worker1(ctx1, "work--1")
	go worker1(ctx1, "work--2")
	time.Sleep(2 * time.Second)
	cancel1()
	fmt.Println("主程序退出")
}
