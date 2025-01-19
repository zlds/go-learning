package _type

import "testing"

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64
	// Go不支持隐式类型转换
	//b = a
	b = int64(a)
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	// &取指针
	aPtr := &a
	// 不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// string是值类型,默认是空字符串
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
