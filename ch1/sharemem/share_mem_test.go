package sharemem

import (
	"sync"
	"testing"
	"time"
)

// 非安全测试
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 安全测试
func TestCounterThreadSafe(t *testing.T) {
	// 使用sync.Mutex保护counter
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			// defer 语句会在函数执行完毕后执行，这里用于释放锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	// 等待退出(类似于Java中的Join,性能高于sleep,因为sleep不知道需要等待多长时间)
	wg.Wait()
	t.Logf("counter = %d", counter)
}
