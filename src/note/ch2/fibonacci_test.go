package ch2

import "testing"

func TestGetElement(t *testing.T) {
	for i := 0; i <= 10; i++ {
		ele := GetElement(i)
		t.Log("index=", i, ",value=", ele)
	}
}
func TestGetElements(t *testing.T) {
	for i := 0; i <= 10; i++ {
		eles := GetElements(i)
		t.Log("n=", i, ",values=", eles)
	}
}
