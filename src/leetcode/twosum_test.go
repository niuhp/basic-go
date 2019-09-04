package leetcode

import (
	"testing"
	"math/rand"
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	for i, a := range nums {
		tmp := target - a
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == tmp {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
func TestTowSum(t *testing.T) {
	len := 30
	nums := make([]int, 0, len)
	for i := 0; i < len; i++ {
		nums = append(nums, rand.Intn(len*3))
	}
	fmt.Println(nums)
	target := len * 2
	result := TwoSum(nums, target)
	fmt.Println("target =", target, "result =", result)
}
