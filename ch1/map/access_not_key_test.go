package _map

import "testing"

func TestAccessNotKey(t *testing.T) {

	// 访问不存在的key返回0值
	m1 := map[int]int{}
	t.Log(m1[1])
	// 如果一个value值是0，这个时候我们无法区分是key不存在还是key对应的值就是0。这个时候就可以通过if判断来区分
	m1[2] = 0
	t.Log(m1[2])

	if v, ok := m1[2]; ok {
		t.Logf("Key 2's value is %d", v)
	} else {
		t.Log("Key 2 is not existing.")
	}
}
