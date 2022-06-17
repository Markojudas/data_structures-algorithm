package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp_val := arr[i]
		left := i - 1

		for left >= 0 {
			if arr[left] > temp_val {
				arr[left+1] = arr[left]
				left--
			} else {
				break
			}
		}
		arr[left+1] = temp_val
	}
}

func main() {

	arr := []int{4, 2, 7, 1, 3}

	fmt.Println("Unsorted Slice")
	fmt.Println(arr)
	fmt.Println()

	fmt.Println("Insertion Sort")
	insertionSort(arr)
	fmt.Println(arr)
	fmt.Println()
}
