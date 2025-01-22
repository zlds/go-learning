package base

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return "汪汪汪"
}

func (c Cat) Speak() string {
	return "喵喵喵"
}

// MakeAnimalSpeak 通过接口实现统一调用
func MakeAnimalSpeak(a Animal) string {
	return a.Speak()
}
