package combination

// Animal 组合
type Animal interface {
	Speaker
	Runner
	Sleeper
}

type Speaker interface {
	Speak() string
}

type Runner interface {
	Run() string
}

type Sleeper interface {
	Sleep() string
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

func (d Dog) Run() string {
	return "狗跑"
}

func (c Cat) Speak() string {
	return "喵喵喵"
}

func (c Cat) Run() string {
	return "猫跑"
}
