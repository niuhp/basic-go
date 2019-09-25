package ch12

import "testing"

func TestGetType(t *testing.T) {
	s0 := GetType(123)
	t.Log(s0)

	s1 := GetType("hello")
	t.Log(s1)

	s2 := GetType(1234.5)
	t.Log(s2)
}
func TestGetType2(t *testing.T) {
	s0 := GetType2(12322)
	t.Log(s0)

	s1 := GetType2("hello22")
	t.Log(s1)

	s2 := GetType2(1234.533)
	t.Log(s2)
}
