package main

import (
	"log"
	"time"
)

func bubbleSort(arr []int) []int {
	start := time.Now()

	size := len(arr)
	var swapped bool

	for i := 0; i < size-1; i++ {
		swapped = false
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	elapsed := time.Since(start)
	log.Printf("Bubble sort %s", elapsed)
	return arr
}
