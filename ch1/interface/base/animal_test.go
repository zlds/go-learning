package base

import "testing"

func TestDogSpeak(t *testing.T) {
	dog := Dog{Name: "柯基"}
	// 调用统一
	speak := MakeAnimalSpeak(dog)
	t.Log(speak)
}

func TestCatSpeak(t *testing.T) {
	cat := Cat{Name: "汤姆"}
	speak := MakeAnimalSpeak(cat)
	t.Log(speak)
}
