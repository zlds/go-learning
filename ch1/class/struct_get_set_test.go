package class

import "testing"

type Person struct {
	Id   int
	Name string
	Age  int
}

// 值接收者(用于获取字段值，但不会修改对象本身)
func (e Person) GetName() string {
	return e.Name
}

// 指针接收者(用于修改结构体的字段值)
func (e *Person) SetName(name string) {
	e.Name = name
}

func TestPerson(t *testing.T) {
	p := Person{1, "Tom", 20}
	t.Log(p.GetName())
	p.SetName("Jack")
	t.Log(p.GetName())
}
