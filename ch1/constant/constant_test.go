package constant

import "testing"

const (
	Monday1    = 1
	Tuesday1   = 2
	Wednesday1 = 3
	Thursday1  = 4
	Friday1    = 5
	Saturday1  = 6
	Sunday1    = 7
)

func TestConstant1(t *testing.T) {
	t.Log(Monday1, Tuesday1)
}
