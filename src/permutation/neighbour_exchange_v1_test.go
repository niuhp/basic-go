package permutation

import (
	"fmt"
	"testing"
)

func TestGetAllPermutationV1(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := GetAllPermutationV1(arr)
	for _, item := range result {
		fmt.Println(item)
	}
}
