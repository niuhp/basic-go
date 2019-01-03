package base

import (
	"testing"
	"fmt"
)

func TestAnimal(t *testing.T) {

	var a1 Animal = Animal{1}
	fmt.Println(a1.age)

	var a2 Dog = Dog{Animal{2}, 4}
	fmt.Println(a2.age)

	var a3 Bird = Bird{Animal{3}, 2}
	fmt.Println(a3.age, "----")
}
