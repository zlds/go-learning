package base

import (
	"errors"
	"fmt"
)

// 定义接口
type UserService interface {
	Add(user UserDTO)
	Get(userId int) (UserVO, error)
}

// UserServiceImpl 实现类(将方法的接收者绑定到该实例上面实现了接口的实现)

type UserServiceImpl struct {
}

// Add 将 Add方法绑定到UserServiceImpl类型上，Add是属于UserServiceImpl的方法。
func (s *UserServiceImpl) Add(user UserDTO) {
	fmt.Printf("添加用户: %v\n", user)
}

// Get 将 Get方法绑定到UserServiceImpl类型上面，Get是属于UserServiceImpl的方法。
func (s *UserServiceImpl) Get(userId int) (UserVO, error) {
	fmt.Printf("获取用户: %v\n", userId)
	// 返回模拟用户数据
	if userId == 1 {
		return UserVO{Id: 1, Name: "Tom"}, nil
	}
	return UserVO{}, errors.New("用户不存在")
}
