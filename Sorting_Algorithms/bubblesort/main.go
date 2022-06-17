package main

import "fmt"

/*
	Bubble Sort is a basic sorting algorithm
	1)	point to two consecutive values in the array. (initially, we start by pointing to the array's first two values)
		Compare the first item with the second

	2)	if the two itmers are out of order, swap them

	3)	move the "pointers" one cell to the right

	4)	repeat steps 1 - 3 until end of array, or if we reach the values that have already been sorted. This completes the first pass-through

	5) 	move pointers back to the first two values of the array and execute another pass-through. Repeat all steps until a pass-thrrough
		with not swaps.

	KEY: at the end of each pass-through the right side of the array is sorted (the largest element is in its place)
*/

func main() {
	arr := []int{65, 55, 45, 35, 25, 15, 10}

	fmt.Println("unsorted array")
	fmt.Println(arr)

	bubbleSort(arr)
	fmt.Println("\nSorted Array")
	fmt.Println(arr)

}

func bubbleSort(arr []int) {
	sorted := false
	lastUnsortedIndex := len(arr) - 1

	numOfsteps := 0
	numOfswap := 0

	for !sorted {
		sorted = true

		//pass-throughs
		for i := 0; i < lastUnsortedIndex; i++ {
			numOfsteps++
			if arr[i] > arr[i+1] {
				numOfswap++
				//swap! since we are doing a swap we have to do another pass through until NO swaps
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}

		lastUnsortedIndex-- //since the last index is sorted after passthrough, no need to visit this cell (-1)
	}

	fmt.Println("Number of comparions:", numOfsteps)
	fmt.Println("Number of swaps:", numOfswap)
}
