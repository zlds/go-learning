package condition

import "testing"

// 多返回值测试
func TestIfMulti(t *testing.T) {
	// if的条件里可以赋值
	if v, err := someFun(); err == nil {
		t.Log(v)
		t.Log("1==1")
	} else {
		t.Log("1!=1")
	}
}

func someFun() (int, error) {
	return 1, nil
}
