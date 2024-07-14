package main

import (
    "fmt"
    "strconv"
)

func main() {
    arr := []int{5,6,7,1,8,9,2,3,10,4}
    for i := 0; i < len(arr)-1; i++ {
        largest := i
        for j := i+1; j < len(arr); j++{
            if arr[j] > arr[largest] {
                largest = j
            }   
        }
        if largest != i {
            arr[largest] += arr[i]
            arr[i] = arr[largest] - arr[i]
            arr[largest] -= arr[i]
        }
    }
    
    var spc string 
    for i := 0; i < len(arr); i++ {
        fmt.Print(spc + strconv.Itoa(arr[i]))
        spc = ", "
    }
    fmt.Println()
}

