package array

import "testing"

/* 切片支持动态扩容。数组和切片声明很像，数组声明时需要指定固定大小，而切片不需要指定大小*/
func TestSlice(t *testing.T) {
	var s0 []int
	// 切片的长度和容量
	t.Log(len(s0), cap(s0))
	// 添加内容
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
}

// 动态扩容(默认空间不够用时会使用上一个大小乘以2)。
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2)
	t.Log(len(Q2), cap(Q2))
}
