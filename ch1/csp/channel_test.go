package csp

import (
	"fmt"
	"testing"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("生产者: 生产 %d\n", i)
		// 发送数据
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	// 关闭channel，不再生产数据
	close(ch)
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Printf("消费者: 消费 %d\n", v)
		time.Sleep(time.Millisecond * 500)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan int, 3)
	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Second * 5)
	fmt.Printf("主程序结束.")
}

type Message struct {
	id   int
	text string
}

// 支持自定义数据类型
func TestChannelType(t *testing.T) {
	// 创建一个channel
	ch := make(chan Message, 2)
	// 存数据
	ch <- Message{1, "hello"}
	ch <- Message{2, "world"}
	// 取数据
	msg1 := <-ch
	msg2 := <-ch
	fmt.Println(msg1.text, msg2.text)
}
