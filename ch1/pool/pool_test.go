package pool

import (
	"fmt"
	"sync"
	"testing"
)

type MyObject struct {
	// 省略
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("创建新对象")
			return "New Object"
		},
	}
	// 从池中获取对象，如果池中没有对象，则调用New函数创建一个新对象
	obj1 := pool.Get().(string)
	fmt.Println("获取obj1对象：", obj1)
	// 将对象放回池中(默认情况下创建对象是不会放到池中的，需要手动调用Put方法)
	pool.Put("Put Object")
	// 再次从池中获取对象
	obj2 := pool.Get().(string)
	fmt.Println("获取obj2对象：", obj2)
	// 再次获取对象(池为空，调用New创建新对象)
	obj3 := pool.Get().(string)
	fmt.Println("获取obj3对象：", obj3)

	// 创建一个对象池
	pool1 := sync.Pool{
		New: func() interface{} {
			fmt.Println("创建新的MyObject")
			return &MyObject{}
		},
	}
	obj4 := pool1.Get().(*MyObject)
	fmt.Println("获取obj4对象：", obj4)
	pool1.Put(&MyObject{})
	obj5 := pool1.Get().(*MyObject)
	fmt.Println("获取obj5对象：", obj5)
}
