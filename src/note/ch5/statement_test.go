package ch5

import (
	"testing"
)

func TestFor(t *testing.T) {
	for i := 1; i <= 20; i += 3 {
		t.Log("i =", i)
	}
}
func TestWhile(t *testing.T) {
	i := 5
	for i < 10 {
		t.Log("i =", i)
		i ++
	}
}
func TestMultiCase(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 3, 6, 9:
			t.Log("i =", i, ",可以被 3 整除")
			//case 0,7: case 中元素重复时会不能编译
		case 4, 8:
			t.Log("i =", i, ",可以被 4 整除")
		default:
			t.Log("i =", i, ",不可以被 3 或 4 整除")
		}
	}
}
func TestConditionCase(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i >= 0 && i < 2:
			t.Log("i =", i, ",i >= 0 && i < 2")
		case i >= 2 && i < 5: // 条件有重复时只命中第 1 个 case
			t.Log("i =", i, ",i >= 2 && i < 5")
		case i >= 3 && i < 6:
			t.Log("i =", i, ",i >= 3 && i < 6")
		default:
			t.Log("i =", i, ",i >= 6")
		}
	}
}
