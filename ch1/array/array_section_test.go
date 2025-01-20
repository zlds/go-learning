package array

import "testing"

func TestArraySection(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	// 获取数组的切片(不包括最后一位)
	arr1Sec := arr[:3]
	t.Log(arr1Sec)
	/* 不支持负数索引 */
	//arr2Sec := arr[:-1]
	//t.Log(arr2Sec)
}
