package permutation

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	arr := []int{1, 2, 3, 4};
	for arr != nil {
		fmt.Println(arr);
		arr = NextPermutation(arr);
	}
}
func TestNextPermutation2(t *testing.T) {
	arr := []int{1, 2, 3, 4,5,6};
	for arr != nil {
		fmt.Println(arr);
		arr = NextPermutation(arr);
	}
}