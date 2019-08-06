package ch06

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	//var s0 []int // 默认 cap 为 1，每次扩容为原来的 2 倍
	//var s0 = make([]int, 5) // 初始化后，len = 5，cap = 5，前 5 个元素都为 0，每次扩容为原来的 2 倍
	var s0 = make([]int, 2, 7) // 初始化后，len = 2，cap = 7，前 2 个元素都为 0，每次扩容为原来的 2 倍
	for i := 0; i < 20; i++ {
		s0 = append(s0, i)
		t.Log("len =", len(s0), ",cap =", cap(s0), ",s0 =", s0)
	}
}

func TestSliceCap(t *testing.T) {
	nums := [16]int{}
	for i := 0; i < 16; i++ {
		nums[i] = i + 1
	}
	t.Log(nums)

	s0 := nums[4:9]
	t.Log("len =", len(s0), "cap =", cap(s0), "s0 =", s0)
	s0[0] = 100 //会改变原数组 nmms[4] 的值
	t.Log(nums)
	s0 = append(s0, 111) //会改变原数组 nmms[9] 的值
	t.Log(nums)
	t.Log("len =", len(s0), "cap =", cap(s0), "s0 =", s0)

	s1 := nums[9:12]
	t.Log("len =", len(s1), "cap =", cap(s1), "s1 =", s1)
}
