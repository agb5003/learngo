package main

import "fmt"

func main() {
	arr := []int{2,4,5,9,8}

	sliced := arr[1:5]

	fmt.Println(sliced)
}