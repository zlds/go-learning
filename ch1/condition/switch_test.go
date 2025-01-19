package condition

import "testing"

// switch支持多条件匹配
func TestSwitchMulti(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("不是0-3之间的数")
		}
	}
}
