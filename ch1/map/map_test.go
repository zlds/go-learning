package _map

import "testing"

func TestInitMap(t *testing.T) {
	// 初始化有值的map
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[1])
	t.Logf("len m1=%d", len(m1))
	// 初始化空map
	m2 := map[int]int{}
	m2[1] = 11
	t.Logf("len m2=%d", len(m2))
	// 通过make创建
	m3 := make(map[int]int, 10)
	m3[1] = 12
	t.Logf("len m3=%d", len(m3))
}
