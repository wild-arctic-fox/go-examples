package customsort

func CocktailSort(a []int) []int {
	arr := make([]int, len(a))
	copy(arr, a)
	swapped := true
	start := 0
	end := len(arr)

	for swapped == true {
		swapped = false
		for i := start + 1; i < end-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swapped = true
			}
		}
		if swapped == false {
			break
		}
		swapped = false
		end = end - 1
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swapped = true
			}
		}
		start = start + 1
	}
	return arr
}
