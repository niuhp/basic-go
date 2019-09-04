package leetcode

import (
	"testing"
	"math/rand"
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return nil
	}
	for i, a := range nums {
		result := []int{a, 0}
		for j, b := range nums {
			if i != j && a+b == target {
				result[1] = b
				return result
			}
		}
	}
	return nil
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
