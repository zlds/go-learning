package _map

import "testing"

func TestForMap(t *testing.T) {

	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Logf("key=%d, value=%d", k, v)
	}

}
