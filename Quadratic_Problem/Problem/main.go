package main

import "fmt"

func main() {
	arr := []int{1, 5, 3, 9, 1, 4}

	hasDuplicate, numOfStepsTaken := hasDuplicateValue(arr)

	fmt.Println("has duplicates:", hasDuplicate)
	fmt.Println("Number of steps taken by the algorithm:", numOfStepsTaken)
	fmt.Println("***********************************")
	fmt.Println("Worst Case: Array of length 5 has no duplicates")

	arr = []int{1, 4, 5, 2, 9} //only 5 elements

	hasDuplicate, numOfStepsTaken = hasDuplicateValue(arr)
	fmt.Println("has duplicates:", hasDuplicate)
	fmt.Println("Number of steps taken by the algorithm:", numOfStepsTaken) //25 steps! O(n^2)

}

func hasDuplicateValue(arr []int) (bool, int) {
	steps := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			steps++
			if i != j && arr[i] == arr[j] {
				return true, steps
			}
		}
	}
	return false, steps
}
