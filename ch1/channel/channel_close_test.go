package channel

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭channel，不再生产数据(告知所有消费者通道关闭)
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// 如果ok为false，说明channel已经关闭(这种方式可以解决多个消费者的情况)
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}

		}
		wg.Done()
	}()
}

func dataReceiver1(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 11; i++ {
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()

}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// 测试读取关闭的channel
	//wg.Add(1)
	//dataReceiver1(ch, &wg)
	wg.Wait()
	fmt.Println("主程序结束.")
}
