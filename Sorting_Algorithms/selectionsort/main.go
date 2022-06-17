package main

import "fmt"

func main() {
	arr := []int{65, 55, 45, 35, 25, 15, 10}

	fmt.Println("Unsorted Slice")
	fmt.Println(arr)
	fmt.Println()

	selectionSort(arr)

	fmt.Println("\nSorted Slice | Selection Sort")
	fmt.Println(arr)
	fmt.Println()
}

func selectionSort(arr []int) {

	numComp := 0
	numSwap := 0

	for i := 0; i < len(arr)-1; i++ {
		smallestThusFarIndex := i

		for j := i + 1; j < len(arr); j++ {
			numComp++
			if arr[j] < arr[smallestThusFarIndex] {
				smallestThusFarIndex = j
			}
		}

		if smallestThusFarIndex != i {
			numSwap++
			arr[i], arr[smallestThusFarIndex] = arr[smallestThusFarIndex], arr[i]
		}
	}

	fmt.Println("Number of comparisions:", numComp)
	fmt.Println("Number of swaps:", numSwap)
}
