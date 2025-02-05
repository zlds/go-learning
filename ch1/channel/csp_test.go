package channel

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service done"
}

func otherTask() {
	fmt.Println("otherTask start")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func asyncService() chan string {
	// 默认情况下必须等待客户端将结果拿走之后才会执行下一步操作
	//retCh := make(chan string)
	// 带缓存的channel(),运行最多1条数据在channel中，而不阻塞下一步操作
	retCh := make(chan string, 1)
	go func() {
		ret := service
		fmt.Println("returned result.")
		// 获取结果(阻塞等待)
		retCh <- ret()
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)
}
