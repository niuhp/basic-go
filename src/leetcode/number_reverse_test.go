package leetcode

import (
	"testing"
	"math"
)

func reverse(x int) int {

	if x < 0 {
		return -reverse(-x)
	} else if x == 0 {
		return 0
	} else {
		tmp := make([]int, 0)
		for x > 0 {
			n := x % 10
			tmp = append(tmp, n)
			x = x / 10
		}

		result := 0
		for index, num := range tmp {
			pow10 := math.Pow10(len(tmp) - index - 1)
			result += num * int(pow10)
		}

		if result>>31 > 0 {
			result = 0
		}

		return result
	}

}

func TestReverse(t *testing.T) {

	a := 123
	aa := reverse(a)
	t.Log(aa)

}
