package main

import (
	"log"
	"time"
)

func quickSort(arr []int) []int {
	size := len(arr) // [2, 2]

	if size <= 1 {
		return arr
	}

	pivot_i := size / 2 // 1

	var left []int
	var right []int

	for i := 0; i < size; i++ {
		if i == pivot_i {
			continue
		}

		if arr[i] < arr[pivot_i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(quickSort(left), append([]int{arr[pivot_i]}, quickSort(right)...)...)
}

func applyQuickSort(arr []int) []int {
	start := time.Now()

	arr = quickSort(arr)

	elapsed := time.Since(start)
	log.Printf("Selection sort %s", elapsed)
	return arr
}
