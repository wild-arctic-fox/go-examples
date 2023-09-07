package customsort

func GnomeSort(a []int) []int {
	arr := make([]int, len(a))
	copy(arr, a)

	for index := 0; index < len(arr); {
		if index == 0 {
			index++
		}
		if arr[index] >= arr[index-1] {
			index++
		} else {
			temp := 0
			temp = arr[index]
			arr[index] = arr[index-1]
			arr[index-1] = temp
			index--
		}
	}
	return arr
}
