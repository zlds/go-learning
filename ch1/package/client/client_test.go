package client

import (
	// 导入依赖包
	"go-learning/ch1/package/order"
	"go-learning/ch1/package/user"
	"testing"
)

func TestClient(t *testing.T) {
	t.Log(order.GetOrder(1))
	t.Log(user.GetUser(1))
}
