package main

import "fmt"

/*
	Page 61, Exercise # 4
	rewrite the function so that it becomes a speedy O(N)

*/

func main() {

	arr := []int{1, 2, 3, 44, 5}

	greatestNum, numSteps := greatestNumber(arr)

	fmt.Println("Greatest Number:", greatestNum)
	fmt.Println("Num of Steps:", numSteps) // 5 O(N)
}

/*
O(n^2) function to write (in python)

def greatestNumber(array):
	for i in array:
		#assume for now that i is the greatest:
		isIValTheGreatest = True

		for j in array:
			#If we find another value that is greater than i,
			#i is not the greatest:
			if j > i:
				isIValTheGreatest = False

		#If, by the time we checked all the other numbers, i
		#is still the greatest, it means that i is the greatest numer:
		if isIValTheGreatest:
			return i
*/

func greatestNumber(arr []int) (gn int, st int) {

	greatNum := arr[0]
	var steps int
	for _, val := range arr {
		steps++
		if val > greatNum {
			greatNum = val
		}
	}

	return greatNum, steps
}
