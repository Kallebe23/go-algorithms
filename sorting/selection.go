package main

import (
	"log"
	"time"
)

func selectSort(arr []int) []int {
	start := time.Now()

	size := len(arr)
	for i := 0; i < size-1; i++ {
		smallest_index := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[smallest_index] {
				smallest_index = j
			}
		}
		temp := arr[i]
		arr[i] = arr[smallest_index]
		arr[smallest_index] = temp
	}

	elapsed := time.Since(start)
	log.Printf("Selection sort %s", elapsed)
	return arr
}
