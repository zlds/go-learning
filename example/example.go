package example

// 导入了fmt包
import "fmt"

/*
*
对外暴露方法
*/
func SayHello() {
	fmt.Println("")
}

/*
*
这是私有方法
*/
func f1() {
	fmt.Print("我是私有方法")
}
