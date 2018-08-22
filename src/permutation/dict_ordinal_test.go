package permutation

import (
	"fmt"
	"testing"
)

func TestNextPermutation1(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	printAllPermutation(arr);
}
func TestNextPermutation2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	printAllPermutation(arr);
}
func printAllPermutation(arr []int) {
	for arr != nil {
		fmt.Println(arr)
		arr = NextPermutation(arr)
	}
}
