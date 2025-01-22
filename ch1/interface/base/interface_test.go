package base

import "testing"

// 定义接口
type Programmer interface {
	WriteHelloWorld() string
}

// 定义结构体
type GoProgrammer struct{}

// 实现接口
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
