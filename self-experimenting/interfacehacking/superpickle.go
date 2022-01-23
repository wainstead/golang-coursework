package main

type first interface {
	funky() string
}
type second interface {
	shunny() string
}

type third interface {
	first
	second
}

type X struct{}
type Y struct{}

func (xer X) funky() string {
	return "I am funky"
}

func (xer X) shunny() string {
	return "I am shunned"
}

func (yer Y) funky() string {
	return "I am funky"
}

func (yer Y) shunny() string {
	return "I am shunned"
}

func Gogo(t third) {
	t.funky()
	t.shunny()
}
