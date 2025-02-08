package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3) = %d; expected %d", result, expected)
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	// Error测试失败会继续执行，其他测试不影响
	t.Error("Error")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	// Fatal测试失败会终止测试，其他测试不影响
	t.Fatal("Error")
	fmt.Println("End")
}

// 断言测试
func TestAddWithAssert(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3), "正确结果是相等的")
}
