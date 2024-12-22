package main

import (
	utils "aoc2024"
	"fmt"
	"os"
)

func main() {
	arr, err := utils.ReadFileIntoArrayOfIntArray("input.txt")

	if err != nil {
		os.Exit(1)
	}

	safeResults := 0;

	for _, outer := range arr {

		safe:= true;

		safe = tolerantRow(outer)

		if (safe) {
			safeResults++
		} else {
			//dampner used - check if the row could be safe
			safe = unsafeCheck(outer)

			if (safe) {
				safeResults++
			}
		}

	}

	fmt.Println(safeResults)

}

func tolerantRow(arr []int) bool {

	safe := true;
	incforTheRow := true;

	//fmt.Println(arr)
	increaseCheck := arr[1] - arr[0];

	if (increaseCheck==0) {
		// the first and second val are the same which is unsafe row
		safe = false;
		return false;
	} else {
		if increaseCheck > 0 {
			incforTheRow = true
		} else {
			incforTheRow = false
		}
	}
	
	for j, inner := range arr {
		if j==0 {
			continue
		}

		tol, inc := tolerant(arr[j-1], inner)
		if tol && (inc == incforTheRow) {
			continue
		} else {
			safe = false
			break;
		}
	}

	return safe;
}

func tolerant(val1 int, val2 int) (bool, bool){
	diff := val1-val2

	if (diff > 3 || diff < -3 || diff == 0){
		 return false, false;
	}

	return true, diff<0;
}

func unsafeCheck(arr []int) bool {

	for i,_ := range arr {
		arr2 := make([]int, 0)

		for j, innerVal := range arr {
			if i != j {
				arr2 = append(arr2, innerVal)
			}
		}
			
		// check for safety, if any safe, return true
		if (tolerantRow(arr2)) {
			return true;
		}
	}
	return false;
}