package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	for true {
		checker := true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				checker = false
			}

		}
		if checker {
			break
		}
	}
	return arr
}

func main() {
	arr := []int{10, 1, 20, 4, 2, 9, 5, 7, 6} // your array
	arr = bubbleSort(arr)

	fmt.Println(arr)
}
