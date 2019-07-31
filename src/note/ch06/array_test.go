package ch06

import "testing"

func TestArrayTravel(t *testing.T) {
	arr := [...]int{11, 22, 33, 44, 55, 66}

	for i := 0; i < len(arr); i++ {
		t.Log("i =", i, ",arr[", i, "] =", arr[i])
	}

	for i, v := range arr {
		t.Log("i =", i, ",v =", v)
	}

	for _, v := range arr {
		t.Log("v =", v)
	}
}
func TestArraySection(t *testing.T) {
	arr := [...]int{11, 22, 33, 44, 55, 66}
	t.Log("arr =", arr)
	arr1 := arr[1:3] //截取数组下标 >=1 且 <3 的元素
	t.Log("arr1 =", arr1)
	arr2 := arr[:3] //截取数组下标 <3 的元素
	t.Log("arr2 =", arr2)
	arr3 := arr[1:] //截取数组下标 >=1 的元素
	t.Log("arr3 =", arr3)
	arr4 := arr[:] //截取所有元素
	t.Log("arr4 =", arr4)
}
