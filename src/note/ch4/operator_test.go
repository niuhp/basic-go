package ch4

import (
	"testing"
	"log"
)

func TestArrayCompare(t *testing.T) {
	a0 := [...]int{1, 2, 3}
	a1 := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 4}
	d := [...]int{1, 3, 2}

	t.Log("a0 == a1", a0 == a1)
	t.Log("b =", b)
	//t.Log("a0 == b",a0 == b) // 元素个数不一样编译不通过
	t.Log("a0 == c", a0 == c)
	t.Log("a0 == d", a0 == d)
}

func TestClearBit(t *testing.T) {
	a := 7
	t.Log(a &^ 1) // 把第 1 位变成 0
	t.Log(a &^ 2) // 把第 2 位变成 0
	t.Log(a &^ 3) // 把第 1、2 位变成 0
	t.Log(a &^ 4) // 把第 3 位变成 0
	t.Log(a &^ 5) // 把第 1、3 位变成 0
	t.Log(a &^ 6) // 把第 2、3 位变成 0
	t.Log(a &^ 7) // 把第 1、2、3 位变成 0
}

func TestAnd(t *testing.T) {
	result := isLeftBigger(9, 10) && isLeftBigger(20, 10) // 同样是短路与
	t.Log(result)
}

func isLeftBigger(a int, b int) bool {
	r := a > b;
	log.Print("a =", a, ", b =", b, ", result =", r)
	return r;
}
