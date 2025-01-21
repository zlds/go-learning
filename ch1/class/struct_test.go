package class

import (
	"fmt"
	"testing"
)

// 定义结构体
type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestEmployee(t *testing.T) {
	// 初始化直接赋值
	e := Employee{"0", "Bob", 20}
	t.Log(e)
	// 通过字段名赋值
	e1 := Employee{Name: "Mike", Age: 30}
	t.Log(e1)
	// new 返回指针
	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22
	t.Logf("e2 is %T", e2)
	t.Log(e2)
}

// 值接收
func (e Employee) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Age: %d", e.Id, e.Name, e.Age)
}
func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e.String())
	e1 := Employee{Name: "Jack", Age: 30}
	e1.UpdateName("Alice")
	t.Log(e1.String())
}
