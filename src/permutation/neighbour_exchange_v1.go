package permutation

func GetAllPermutationV1(arr []int) [][]int {
	len := len(arr)
	if len == 1 {
		return [][]int{arr}
	}
	lastValue := arr[len-1]
	subArr := arr[:len-1]
	subArrs := GetAllPermutationV1(subArr)
	var result [][]int
	for _, item := range subArrs {
		subResult := getPermutations(item, lastValue)
		result = append(result, subResult...)
	}
	return result
}
func getPermutations(arr []int, value int) [][]int {
	len := len(arr) + 1
	var result [][]int
	for i := len - 1; i >= 0; i-- {
		var tmp []int
		tmp = append(tmp, arr[:i]...)
		tmp = append(tmp, value)
		tmp = append(tmp, arr[i:]...)
		result = append(result, tmp)
	}
	return result
}
