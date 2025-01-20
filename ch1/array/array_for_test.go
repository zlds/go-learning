package array

import "testing"

func TestArrayFor(t *testing.T) {
	// 基础遍历
	t.Log("基础遍历")
	var arr = [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	// for range遍历
	t.Log("for range遍历")
	for idx, e := range arr {
		t.Log(idx, e)
	}
	// 不需要索引(-符号代表占位符)
	t.Log("不需要索引")
	for _, e := range arr {
		t.Log(e)
	}
}
