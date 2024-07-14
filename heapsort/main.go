// Simple heapsort implementation in Go
package main

import (
	"fmt"
	"strconv"
)

func heapify(arr []int, n, i int) {
    smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}

	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]

		heapify(arr, n, smallest)
	}
}

func heapsort(arr []int) {
	// Build min heap
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}

	// Remove elements from the top
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{2, 5, 7, 6, 8, 9, 3, 4, 1, 10}
	heapsort(arr)

	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(strconv.Itoa(arr[i]) + ", ")
	}
	fmt.Print(arr[len(arr)-1])
    fmt.Println()
}
