package main

import (
	"log"
	"time"
)

func recursiveSplitMerge(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	half := len(arr) / 2
	left := arr[:half]
	right := arr[half:]

	splitedLeft := recursiveSplitMerge(left)
	splitedRight := recursiveSplitMerge(right)

	return merge(splitedLeft, splitedRight)
}

func merge(arr1 []int, arr2 []int) []int {
	var newArr []int

	var i = 0
	var j = 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			newArr = append(newArr, arr1[i])
			i += 1
		} else {
			newArr = append(newArr, arr2[j])
			j += 1
		}
	}

	for ; i < len(arr1); i++ {
		newArr = append(newArr, arr1[i])
	}

	for ; j < len(arr2); j++ {
		newArr = append(newArr, arr2[j])
	}

	return newArr
}

func mergeSort(arr []int) []int {
	start := time.Now()

	arr = recursiveSplitMerge(arr)

	elapsed := time.Since(start)
	log.Printf("Merge sort %s", elapsed)
	return arr
}
