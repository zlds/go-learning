package singleton

import (
	"fmt"
	"sync"
	"testing"
)

type Singleton struct{}

var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("创建 Singleton实例")
		// 使用指针类型确保所有调用都返回相同的实例
		instance = &Singleton{}
	})
	return instance
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			GetInstance()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("主程序结束.")
}
