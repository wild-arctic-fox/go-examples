package main

import (
	"example/sorting/generators"
	"example/sorting/sort"
	"fmt"
	"reflect"
	"sort"
	"time"
)

func isSorted(arr []int) bool {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	return reflect.DeepEqual(arr, sortedArr)
}

func main() {
	fmt.Println("Sort comparasion")

	for n := 1000; n < 1_000_000; n *= 5 {
		// random arr
		arr := randomint.GenerateRandomInt(n)
		initRes := isSorted(arr)
		fmt.Println("Initial array sorted:", initRes)
		fmt.Println("Elements:", n)

		// bubble sort
		start := time.Now()
		bubbleSorted := customsort.BubbleSort(arr)
		bubbleDuration := time.Since(start).Milliseconds()
		bubbleRes := isSorted(bubbleSorted)
		fmt.Println("\nBubble sort time:", bubbleDuration, "\nCorrectly sorted:", bubbleRes)

		// insert sort
		start = time.Now()
		insSorted := customsort.InsertionSort(arr)
		insertDuration := time.Since(start).Milliseconds()
		insRes := isSorted(insSorted)
		fmt.Println("\nInsert sort time:", insertDuration, "\nCorrectly sorted:", insRes)

		// cocktail sort
		start = time.Now()
		cocktailSorted := customsort.CocktailSort(arr)
		cocktailDuration := time.Since(start).Milliseconds()
		cocktailRes := isSorted(cocktailSorted)
		fmt.Println("\nCocktail sort time:", cocktailDuration, "\nCorrectly sorted:", cocktailRes)

		// gnome sort
		start = time.Now()
		gnomeSorted := customsort.GnomeSort(arr)
		gnomeDuration := time.Since(start).Milliseconds()
		gnomeRes := isSorted(gnomeSorted)
		fmt.Println("\nGnome sort time:", gnomeDuration, "\nCorrectly sorted:", gnomeRes)

		modifiableArr := make([]int, len(arr))
		copy(modifiableArr, arr)
		// merge sort
		start = time.Now()
		mergeSorted := customsort.MergeSort(modifiableArr, 0, len(modifiableArr)-1)
		mergeDuration := time.Since(start).Milliseconds()
		mergeRes := isSorted(mergeSorted)
		fmt.Println("\nMerge sort time:", mergeDuration, "\nCorrectly sorted:", mergeRes)

		// selection sort
		start = time.Now()
		selectionSorted := customsort.SelectionSort(arr)
		selectionDuration := time.Since(start).Milliseconds()
		selectionRes := isSorted(selectionSorted)
		fmt.Println("\nSelection sort time:", selectionDuration, "\nCorrectly sorted:", selectionRes)
	}
}
