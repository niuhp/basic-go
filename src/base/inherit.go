package base

import "fmt"

type Animal struct {
	age int
}

type Dog struct {
	Animal
	leg int
}

type Bird struct {
	Animal
	wing int
}

func (a Animal) SayAge() {
	fmt.Println("my age is ", a.age)
}
func (a Animal) SayAge1() {
	fmt.Println("my age is ", a.age)
}