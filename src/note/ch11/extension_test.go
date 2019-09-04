package ch11

import (
	"fmt"
	"testing"
)

type People struct {
}

func (p *People) Greet() {
	fmt.Println("hello...")
}

func (p *People) GreetTo(name string) {
	p.Greet()
	fmt.Println(name)
}

type Man struct {
	People
}

func (m *Man) Greet() {
	fmt.Println("nihao!")
}

func TestGreet(t *testing.T) {
	m := new(Man)

	m.GreetTo("tiny") // 可以直接调用 People 定义的函数，但不能覆盖

	m.Greet()
	m.People.Greet()

	fmt.Println("-------------------")

	var m2 *Man = new(Man)
	m2.Greet()
	//var p *People = new(Man) //不能赋值给父类

}
