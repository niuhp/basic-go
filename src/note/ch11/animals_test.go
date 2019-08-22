package ch11

import "testing"

func TestCat_SayHello(t *testing.T) {
	var a Animal
	a = new(Cat)
	a.SayHello()
}

func TestDog_SayHello(t *testing.T) {
	var a Animal
	a = new(Dog)
	a.SayHello()
}
