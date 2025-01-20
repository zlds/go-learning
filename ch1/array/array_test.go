package array

import "testing"

func TestArray(t *testing.T) {
	// 声明数组
	var arr [3]int
	t.Log(arr)

	// 声明并初始化
	arr1 := [4]int{1, 2, 3, 4}
	t.Log(arr1)

	// 不指定长度
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr2)

	// 多维数组
	grid := [2][2]int{{1, 2}, {3, 4}}
	t.Log(grid)
}
