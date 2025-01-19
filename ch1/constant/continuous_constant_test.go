package constant

import "testing"

// 连续常量
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
}
