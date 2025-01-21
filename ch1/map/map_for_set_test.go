package _map

import "testing"

func TestMapForSet(t *testing.T) {
	// 通过map实现Set功能(map的key是唯一的，value可以是任意类型。如果我们只关心key，可以将value设置为bool类型)
	mySet := map[int]bool{}
	mySet[1] = true
	n := 2
	if mySet[n] {
		t.Logf("%d is existing.", n)
	} else {
		t.Logf("%d is not existing.", n)
	}
	mySet[1] = false
	t.Log(mySet)
}
