package leetcode

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	maxLen, left := 0, 0 // 将最大长度和滑动窗口左边索引初始化为 0
	m := map[int32]int{}
	for i, c := range s {
		if v, ok := m[c]; ok {
			if v > left {
				left = v
			}
		}
		m[c] = i + 1 // map 里保存的是重复字符右边的字符的索引
		currentLen := i - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "   "
	len := lengthOfLongestSubstring(s)
	t.Log("s =", s, ", len =", len)
}
