package main

import (
	"fmt"
	"strconv"
)

func quicksort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := len(arr)-1
    i := -1
    for j := 0; j < pivot; j++ {
        if arr[j] < arr[pivot] {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    i++
    arr[i], arr[pivot] = arr[pivot], arr[i]

    left := quicksort(arr[:i])
    right := quicksort(arr[i+1:])
    merged := append(left, arr[i])
    merged = append(merged, right...)

    return merged
}

func main() {
    arr := []int{2, 4, 5, 7, 6, 3, 8, 9, 10, 1}
    quicksort(arr)

    for i := 0; i < len(arr)-1; i++ {
        fmt.Print(strconv.Itoa(arr[i]) + ", ")
    }
    fmt.Print(arr[len(arr)-1]); fmt.Println()
}
