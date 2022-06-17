package main

import "fmt"

func main() {
	fmt.Println("linear search of an ordered array")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("What number do you wish to find: ")
	var elem int
	_, err := fmt.Scanf("%d", &elem)

	if err != nil {
		fmt.Println(err)
	} else {
		isFound, idx := linearSearch(arr, elem)

		if isFound {
			fmt.Println("Element", elem, "was found in index", idx)
		} else {
			fmt.Println("Element", elem, "was not found")
		}
	}

}

func linearSearch(arr []int, elem int) (bool, int) {

	for idx, val := range arr {
		if val == elem {
			return true, idx
		}
	}

	return false, -1
}
