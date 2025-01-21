package str

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	// 默认是空值
	var s string
	t.Log(s)
	// 赋值并打印长度
	s = "hello"
	t.Log(len(s))
	// rune获取字符串的unicode
	s = "中"
	c := []rune(s)
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "大明王朝"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, parts := range parts {
		t.Log(parts)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log(s)

}
