package customsort

func swap(arr []int, xp int, yp int) {
	temp := arr[xp]
	arr[xp] = arr[yp]
	arr[yp] = temp
}

func SelectionSort(a []int) []int {
	arr := make([]int, len(a))
	copy(arr, a)

	var min_idx int
	for i := 0; i < len(arr)-1; i++ {
		min_idx = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		swap(arr, min_idx, i)
	}

	return arr
}
