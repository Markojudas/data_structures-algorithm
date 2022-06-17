package main

import "fmt"

func main() {
	fmt.Println("Binary Search")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var elem int
	fmt.Print("Select number to find: ")

	_, err := fmt.Scanf("%d", &elem)

	if err != nil {
		fmt.Println("error:", err)
	} else {

		fmt.Println("\nThe Recursive Way")
		isFound, idx := recursiveBS(arr, elem)

		printResult(isFound, elem, idx)

		fmt.Println("\nUsing For loop Way: Binary Search")

		isFound, idx = binarySearch(arr, elem)
		printResult(isFound, elem, idx)

	}

}

func printResult(isFound bool, elem int, idx int) {

	if isFound {
		fmt.Println("Element", elem, "was found at index:", idx)
	} else {
		fmt.Println("Element", elem, "was NOT found")
	}

}

func setLowAndHigh(arr []int) (low int, high int) {

	return 0, len(arr) - 1
}

func binarySearch(arr []int, elem int) (bool, int) {

	low, high := setLowAndHigh(arr)

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == elem {
			return true, mid
		} else if arr[mid] < elem {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false, -1

}

func recursiveBS(arr []int, elem int) (bool, int) {

	low, high := setLowAndHigh(arr)
	return binarySearchRec(arr, elem, low, high)
}

func binarySearchRec(arr []int, elem int, low int, high int) (bool, int) {

	if low > high {
		return false, -1
	}

	mid := low + (high-low)/2

	if arr[mid] == elem {
		return true, mid
	} else if arr[mid] < elem {
		return binarySearchRec(arr, elem, mid+1, high)
	} else {
		return binarySearchRec(arr, elem, low, mid-1)
	}
}
