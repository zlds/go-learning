package operator

import "testing"

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == d)
	// 长度不相等编译出错
	//c := [...]int{1, 2, 3, 4, 5}
	//t.Log(a == c)
}
