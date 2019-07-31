package ch02

import "testing"

func TestGetElement(t *testing.T) {
	for i := 0; i <= 10; i++ {
		ele := GetElement(i)
		t.Log("index =", i, ",value =", ele)
	}
}
func TestGetElements(t *testing.T) {
	for i := 0; i <= 10; i++ {
		eles := GetElements(i)
		t.Log("n =", i, ",values =", eles)
	}
}

func TestChange(t *testing.T) {
	var (
		a = 1
		b = 2
		c = 3
	)
	t.Log("a =", a, ",b =", b)
	a, b, c = b, a, c+1
	t.Log("after change,a =", a, ",b =", b, ",c =", c)
}
