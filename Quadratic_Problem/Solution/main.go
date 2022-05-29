package main

import (
	"fmt"
)

//A LINEAR SOLUTION

func main() {
	arr := []int{1, 4, 5, 2, 9}

	hasDuplicate, numOfSteps := hasDuplicateValue(arr)

	fmt.Println("Has duplicate:", hasDuplicate)
	fmt.Println("Number of steps taken:", numOfSteps)
}

func hasDuplicateValue(arr []int) (bool, int) {

	existingNUmbers := make([]int, 250) //helper array - had to create an array big enough.

	var steps int
	for _, val := range arr {
		steps++
		if existingNUmbers[val] == 1 {
			return true, steps
		} else {
			existingNUmbers[val] = 1
		}
	}
	return false, steps
}
