package ch10

import (
	"fmt"
	"unsafe"
	"testing"
)

type Person struct {
	name string
	age  int
}

func (p Person) PrintInfo1() { //对象的拷贝
	fmt.Printf("address:[%x,%x,%x],type:%T,name:%s,age:%d\n", unsafe.Pointer(&p), unsafe.Pointer(&p.name), unsafe.Pointer(&p.age), p, p.name, p.age)
}
func (p *Person) PrintInfo2() { //对象指针的拷贝
	fmt.Printf("address:[%x,%x,%x],type:%T,name:%s,age:%d\n", unsafe.Pointer(&p), unsafe.Pointer(&p.name), unsafe.Pointer(&p.age), p, p.name, p.age)
}

func TestStruct1(t *testing.T) {
	p := Person{name: "hello", age: 20}
	t.Logf("address:[%x,%x,%x],type:%T,name:%s,age:%d\n", unsafe.Pointer(&p), unsafe.Pointer(&p.name), unsafe.Pointer(&p.age), p, p.name, p.age)
	p.PrintInfo1()
	p.PrintInfo2()

}
func TestStruct2(t *testing.T) {
	p := new(Person)
	p.name = "hi"
	p.age = 100

	t.Logf("address:[%x,%x,%x],type:%T,name:%s,age:%d\n", unsafe.Pointer(&p), unsafe.Pointer(&p.name), unsafe.Pointer(&p.age), p, p.name, p.age)
	p.PrintInfo1()
	p.PrintInfo2()
}
