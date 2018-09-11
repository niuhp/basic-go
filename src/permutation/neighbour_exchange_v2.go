package permutation

type DirectionalInt struct {
	value     int
	direction int
}

func GetAllPermutationV2(arr []int) [][]int {
	var directionalInts []DirectionalInt
	for i := range arr {
		value := arr[i]
		directionalInts = append(directionalInts, DirectionalInt{value, 0})
	}
	var result [][]int
	for directionalInts != nil {
		var tmp []int
		for _, item := range directionalInts {
			tmp = append(tmp, item.value)
		}
		result = append(result, tmp)
		directionalInts = nextPermutation(directionalInts)
	}

	return result
}
func nextPermutation(arr []DirectionalInt) []DirectionalInt {
	maxActiveIndex, maxActive := findMaxActive(arr)
	if maxActiveIndex == -1 {
		return nil
	}
	if maxActive.direction == 0 {
		swap(arr, maxActiveIndex, maxActiveIndex-1)
	} else if maxActive.direction == 1 {
		swap(arr, maxActiveIndex, maxActiveIndex+1)
	}

	for i, current := range arr {
		if current.value > maxActive.value {
			arr[i].direction = current.direction ^ 1
		}
	}
	return arr
}

func findMaxActive(arr []DirectionalInt) (int, DirectionalInt) {

	maxActiveIndex := -1
	var maxActive DirectionalInt
	for i, current := range arr {
		if current.direction == 0 && i > 0 {
			pre := arr[i-1]
			if pre.value < current.value && maxActive.value < current.value {
				maxActiveIndex = i
				maxActive = current
			}
		}
		if current.direction == 1 && i < len(arr)-1 {
			next := arr[i+1]
			if next.value < current.value && maxActive.value < current.value {
				maxActiveIndex = i
				maxActive = current
			}
		}
	}

	return maxActiveIndex, maxActive
}
func swap(arr []DirectionalInt, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
