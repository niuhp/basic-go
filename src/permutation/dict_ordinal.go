package permutation

func NextPermutation(arr []int) []int {
	lastPositive := findLastPositive(arr)
	if lastPositive == -1 {
		return nil
	}
	len := len(arr)
	result := make([]int, len)
	for i := 0; i < lastPositive; i++ {
		result[i] = arr[i]
	}
	lastLarge := findLastLarge(arr, lastPositive)
	result[lastPositive] = arr[lastLarge]
	for i := lastPositive + 1; i < len; i++ {
		if i == len+lastPositive-lastLarge {
			result[i] = arr[lastPositive]
		} else {
			result[i] = arr[len+lastPositive-i]
		}
	}
	return result
}
func findLastPositive(arr []int) int {
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			return i
		}
	}
	return -1
}
func findLastLarge(arr []int, fromIndex int) int {
	for i := len(arr) - 1; i > fromIndex; i-- {
		if arr[i] > arr[fromIndex] {
			return i
		}
	}
	return -1
}
