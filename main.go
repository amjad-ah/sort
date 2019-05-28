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

func selectionSort(arr []int) []int {
	var smallest int
	for i, v := range arr {
		smallest = v
		for j := i; j < len(arr); j++ {
			if arr[j] < smallest {
				smallest, arr[j] = arr[j], smallest
			}
		}
		arr[i] = smallest
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		prev := i - 1

		for prev >= 0 && key < arr[prev] {
			arr[prev+1] = arr[prev]
			prev = prev - 1
		}
	}
	return arr
}

func main() {
	arr := []int{10, 1, 20, 4, 2, 9, 5, 7, 6} // your array
	selection := selectionSort(arr)
	bubble := bubbleSort(arr)
	insertion := insertionSort(arr)

	fmt.Printf("Selection sort: %d\n", selection)
	fmt.Printf("Bubble sort: %d\n", bubble)
	fmt.Printf("Insertion sort: %d\n", insertion)
}
