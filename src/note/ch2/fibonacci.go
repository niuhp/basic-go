package ch2

func GetElement(n int) int {
	if n <= 0 {
		return -1
	} else if n <= 2 {
		return 1
	} else {
		return GetElement(n-1) + GetElement(n-2)
	}
}
func GetElements(n int) []int {
	if n <= 0 {
		return make([]int, 0)
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			result[i] = 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}
	return result
}
