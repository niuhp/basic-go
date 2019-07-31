package ch03

import (
	"testing"
	"math"
)

func TestConvert(t *testing.T) {
	var a int8 = 2
	var b int16 = 4
	var c int64 = 8
	t.Log("a =", a, ",b =", b, ",c =", c)
	// b = a // 编译错误，不能向大的类型转换
	// b = c // 编译错误，不能向小的类型转换
	b = int16(a) //可以这样转
	t.Log("a =", a, ",b =", b, ",c =", c)
	b = int16(c) //可以这样转
	//a = int8(128)// 编译错误，不能超过 127
	b = int16(130) //可以这样转
	a = int8(b)    // 超过 127 ,转换后的值不是期望的
	t.Log("a =", a, ",b =", b, ",c =", c)
	a = int8(127) //可以这样转
	t.Log("a =", a, ",b =", b, ",c =", c)

}

func TestMax(t *testing.T) {
	t.Log("max int64 =", math.MaxInt64) // math 中有这些常量
}

func TestPoint(t *testing.T) {
	a := 2
	ap := &a
	// ap = ap + 1 // 不支持指针运算
	t.Log("a =", a, "ap =", ap)
	t.Logf("type a is:%T,type ap is:%T", a, ap)
}
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log("len(s) =", len(s))
	// t.Log(s == nil) // 没有 nil
	t.Log(s == "")
}
