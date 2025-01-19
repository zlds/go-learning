package condition

import "testing"

func TestIf(t *testing.T) {
	// 支持变量赋值
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}
