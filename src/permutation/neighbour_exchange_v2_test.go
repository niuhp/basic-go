package permutation

import (
	"testing"
	"fmt"
)

func TestGetAllPermutationV2(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := GetAllPermutationV2(arr)
	for i := range result {
		fmt.Println(result[i])
	}
}
