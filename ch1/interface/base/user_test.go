package base

import "testing"

func TestUserServiceImpl_Add(t *testing.T) {
	userService := UserServiceImpl{}
	user := UserDTO{Id: 1, Name: "Tom", Password: "123456"}
	userService.Add(user)
}

func TestUserServiceImpl_Get(t *testing.T) {
	userService := UserServiceImpl{}
	user, err := userService.Get(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}
