package main

import (
	"fmt"
)

func cloneArray(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

func main() {
	const arraySize = 50000
	fmt.Printf("Sorting algorithms - Random list with %v integers\n", arraySize)

	var random = generateRandomIntArr(arraySize)

	// fmt.Println("--- Bubble Sort ---")
	// res := bubbleSort(cloneArray(random))
	// // fmt.Println(res)

	// fmt.Println("--- Selection Sort ---")
	// res = selectSort(cloneArray(random))
	// // fmt.Println(res)

	// fmt.Println("--- Insertion Sort ---")
	// res = insertionSort(cloneArray(random))
	// // fmt.Println(res)

	fmt.Println("--- Merge Sort ---")
	res := mergeSort(cloneArray(random))
	// fmt.Println(res)

	fmt.Println("--- Quick Sort ---")
	res = applyQuickSort(random)
	// fmt.Println(res)

	fmt.Println(len(res))
}
