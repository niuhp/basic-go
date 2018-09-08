package permutation

import (
	"fmt"
	"testing"
)

func TestGetAllPermutation(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := GetAllPermutation(arr)
	for i := range result {
		fmt.Println(result[i])
	}
}
