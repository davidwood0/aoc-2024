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

		safe := true;
		incforTheRow := true;

		increaseCheck := outer[1] - outer[0];

		if (increaseCheck==0) {
			// the first and second val are the same which is unsafe row
			safe = false;
			continue;
		} else {
			if increaseCheck > 0 {
				incforTheRow = true
			} else {
				incforTheRow = false
			}
		}

		for j, inner := range outer {
			if j==0 {
				continue
			}

			tol, inc := tolerant(outer[j-1], inner)
			if tol && (inc == incforTheRow) {
				continue
			} else {

				safe = false
				break;
			}
		}

		if (safe) {
			safeResults++
		}

	}

	fmt.Println(safeResults)

}

func tolerant(val1 int, val2 int) (bool, bool){
	diff := val1-val2

	if (diff > 3 || diff < -3 || diff == 0){
		 return false, false;
	}

	return true, diff<0;
}