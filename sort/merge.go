package customsort

// Merges two subarrays of arr[].
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]
func merge(arr []int, l int, m int, r int) []int {
	n1 := m - l + 1
	n2 := r - m

	// Create temp arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temp arrays L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// Merge the temp arrays back into arr[l..r]
	// Initial index of first subarray
	i := 0
	j := 0
	k := l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of
	// L[], if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of
	// R[], if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
	return arr
}

func MergeSort(arr []int, l int, r int) []int {
	if l >= r {
		return arr
	}
	m := l + (r-l)/2
	arr = MergeSort(arr, l, m)
	arr = MergeSort(arr, m+1, r)
	arr = merge(arr, l, m, r)
	return arr
}
