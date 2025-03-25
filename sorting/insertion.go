package main

import (
	"log"
	"time"
)

func insertionSort(arr []int) []int {
	start := time.Now()

	size := len(arr)

	for i := 1; i < size; i++ {
		temp := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > temp {
				temp2 := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp2
			} else {
				break
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("Insertion sort %s", elapsed)
	return arr
}
