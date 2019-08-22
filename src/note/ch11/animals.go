package ch11

import "fmt"

type Animal interface {
	SayHello()
}

type Cat struct {
}
type Dog struct {
}

func (c *Cat) SayHello() {
	fmt.Println("i am a cat,miao,miao...")
}

func (d *Dog) SayHello() {
	fmt.Println("i am a dog,wang,wang...")
}
