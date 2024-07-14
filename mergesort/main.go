package main

import (
	"fmt"
	"strconv"
)

func mergesort(arr []int) []int {
	if len(arr) >= 2 {
		var merged []int
		left := arr[0:len(arr)/2]
		right := arr[len(arr)/2:]
		left = mergesort(left)
		right = mergesort(right)
		i := 0
		j := 0

		// Merge left and right arrays
		for i < len(left) && j < len(right) {
			if left[i] > right[j] {
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[j])
				j++
			}
		}

		// Copy remaining elements
		for i < len(left) {
			merged = append(merged, left[i])
			i++
		}

		for j < len(right) {
			merged = append(merged, right[j])
			j++
		}

		return merged
	} else {return arr}
}

func main() {
	arr := []int{2, 5, 7, 6, 8, 9, 3, 4, 1, 10}
	arr = mergesort(arr)

	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(strconv.Itoa(arr[i]) + ", ")
	}
	fmt.Print(arr[len(arr)-1]); fmt.Println()
}
