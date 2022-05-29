package main

import "fmt"

/*
	Use Big O to describe the time complexity of the following function.

*/

func main() {

	arr := []int{1, 2, 3, 4, 5} // five elements = O(n^2) = 25 steps

	greatestProd, NumSteps := greatestProduct(arr)

	fmt.Println("Greatest Product of any pair of two numbers within the slice:", greatestProd)
	fmt.Println("Number of steps taken that it took:", NumSteps)

}

func greatestProduct(arr []int) (int, int) {

	greatestProductSoFar := arr[0] * arr[1]
	var steps int

	for i, iVal := range arr {
		for j, jVal := range arr {
			steps++
			prod := iVal * jVal
			if i != j && prod > greatestProductSoFar {
				greatestProductSoFar = prod
			}
		}
	}

	return greatestProductSoFar, steps
}
