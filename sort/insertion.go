package customsort

func InsertionSort(a []int) []int {
	arr := make([]int, len(a))
	copy(arr, a)
	var key, j int
	for i := 0; i < len(arr); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}
