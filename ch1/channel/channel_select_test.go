package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectMultiple(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "来自ch1的消息"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "来自ch2的消息"
	}()

	// 通过select监听多个channel
	select {
	case msg1 := <-ch1:
		fmt.Println("收到：", msg1)
	case msg2 := <-ch2:
		fmt.Println("收到：", msg2)
	}
	// 等待goroutine执行完毕
	time.Sleep(3 * time.Second)
}

func queryDB(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "数据库查询结果"
}

// 超时处理，防止goroutine无限阻塞
func TestChannelTimeOut(t *testing.T) {
	ch := make(chan string)
	go queryDB(ch)
	select {
	case result := <-ch:
		fmt.Println("查询成功：", result)
	case <-time.After(2 * time.Second):
		fmt.Println("查询超时")
	}

}
