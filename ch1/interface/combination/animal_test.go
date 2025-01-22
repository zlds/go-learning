package combination

import "testing"

func TestCat_Run(t *testing.T) {
	cat := Cat{Name: "汤姆"}
	t.Log(cat.Run())
}

func TestCat_Speak(t *testing.T) {
	cat := Cat{Name: "汤姆"}
	t.Log(cat.Speak())
}

func TestDog_Run(t *testing.T) {
	dog := Dog{Name: "柯基"}
	t.Log(dog.Run())
}

func TestDog_Speak(t *testing.T) {
	dog := Dog{Name: "柯基"}
	t.Log(dog.Speak())
}
