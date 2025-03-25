package main

import "math/rand/v2"

func generateRandomIntArr(size int) []int {
	var arr []int

	for len(arr) < size {
		arr = append(arr, rand.IntN(20000))
	}

	return arr
}
