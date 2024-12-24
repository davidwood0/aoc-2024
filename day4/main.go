package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var nextLetter map[string]string = map[string]string{
	"X": "M",
	"M": "A",
	"A": "S",
}

var prevLetter map[string]string = map[string]string{
	"S": "A",
	"A": "M",
	"M": "X",
}


var rowMap = make(map[string]map[int]map[int]bool)

func main() {

	letterCount, err := mapRows("input.txt")

	if err != nil {
		os.Exit(1)
	}

	// now solve with rows, for the least found of X and S
	lowest := letterCount["X"];
	lowestString := "X"

	for i,v:= range letterCount {
		if v < lowest {
			lowest = v;
			lowestString = i;
		}
	}

	println("lowest: ", lowestString)

	result := solveForLetter(lowestString);

	println(result)

}


func mapRows(inFile string) (map[string]int, error) {
	f, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r := bufio.NewReader(f)

	rowCount := 0

	letterCount := make(map[string]int)

	for {
		line, err := r.ReadString('\n')
		// EOF is treated as an error, but we still have last line of data left to process
		// so don't break on that
		// break on any other un-expected errors though
		if err != nil && err != io.EOF {
			break
		}

		split := strings.Split(line, "")

		colCount := 0

		letterColMap := map[string]map[int]bool{}

		for _, v := range split {

			validLetter:= false

			switch (v){
				case "X":
					letterCount[v]++
					validLetter = true
				case "M":
					validLetter = true
				case "A": 
					validLetter = true
				case "S": {
					letterCount[v]++
					validLetter = true;
				} 
			}

			if (validLetter) {

				val, ok := letterColMap[v] 
				if !ok {
					letterColMap[v] = map[int]bool{colCount: true}
				} else {
					val[colCount] = true;
				}
			}
			colCount++;
		}
		
		for i,v := range letterColMap {
			val, ok := rowMap[i] 
			if !ok {
				rowMap[i] = map[int]map[int]bool{rowCount: v};
			} else {
				val[rowCount] = v;
			}
		}

		rowCount++

		// break if we hit EOF
		if err == io.EOF {
			break
		}
	}

	defer f.Close()

	return letterCount, nil
}

func solveForLetter(letter string) int {

	result := 0;
	resultMap := map[string]int{};

	mapToUse := nextLetter;

	if (letter == "S") {
		mapToUse = prevLetter;
	}

	letterMap := rowMap[letter];

	for row, colMap := range letterMap {
		for col := range colMap {
			northCol:= searchDirection(findNorth, mapToUse, letter, row, col)
			northEastDiag:= searchDirection(findNorthEast,  mapToUse, letter, row, col)
			northWestDiag:= searchDirection(findNorthWest,  mapToUse, letter, row, col)
			eastRow:= searchDirection(findEast,  mapToUse, letter, row, col)
			westRow:= searchDirection(findWest,  mapToUse, letter, row, col)
			southCol:= searchDirection(findSouth,  mapToUse, letter, row, col)
			southEastDiag:= searchDirection(findSouthEast,  mapToUse, letter, row, col)
			southWestDiag:= searchDirection(findSouthWest,  mapToUse, letter, row, col)

			if (northCol){ resultMap["N"]++; result++ };
			if (northEastDiag){  resultMap["NE"]++; result++ };
			if (northWestDiag){  resultMap["NW"]++; result++ };
			if (eastRow){  resultMap["E"]++; result++ };
			if (westRow){  resultMap["W"]++; result++ };
			if (southCol){  resultMap["S"]++; result++ };
			if (southEastDiag){  resultMap["SE"]++; result++ };
			if (southWestDiag){  resultMap["SW"]++; result++ };
			
			}
	}

	for i,v := range resultMap {
		println(fmt.Sprintf("%v results: %v", i, v))
	}

	return result;
}

type searchDir func(valToFind string, row int, col int, magnitude int) bool

func searchDirection(findFn searchDir, letterMap map[string]string, startingLetter string, row int, col int) bool {

	foundWord := true;

	magnitude := 1;
	for letterMap[startingLetter] != "" && foundWord {
		foundWord = findFn(letterMap[startingLetter], row, col, magnitude)
		startingLetter = letterMap[startingLetter]
		magnitude++;
	}

	return foundWord
}

func findChar(valToFind string, rowToLook int, colToLook int) bool {
	found := rowMap[valToFind][rowToLook][colToLook];

	return found;
}

func findNorthWest(valToFind string, row int, col int, magnitude int) bool {
	rowToLook := row-magnitude
	colToLook := col-magnitude
	return findChar(valToFind, rowToLook, colToLook)
}

func findNorthEast(valToFind string, row int, col int, magnitude int) bool {
	rowToLook := row-magnitude
	colToLook := col+magnitude
	return findChar(valToFind, rowToLook, colToLook)
}

func findNorth(valToFind string, row int, col int, magnitude int) bool {
	rowToLook := row-magnitude
	return findChar(valToFind, rowToLook, col)
}

func findSouth(valToFind string, row int, col int, magnitude int) bool {
	rowToLook := row+magnitude
	return findChar(valToFind, rowToLook, col)
}

func findSouthWest(valToFind string, row int, col int, magnitude int) bool {
	rowToLook := row+magnitude
	colToLook := col-magnitude
	return findChar(valToFind, rowToLook, colToLook)
}

func findSouthEast(valToFind string, row int, col int, magnitude int) bool {
	rowToLook := row+magnitude
	colToLook := col+magnitude
	return findChar(valToFind, rowToLook, colToLook)
}

func findWest(valToFind string, row int, col int, magnitude int) bool {
	colToLook := col-magnitude
	return findChar(valToFind, row, colToLook)
}

func findEast(valToFind string, row int, col int, magnitude int) bool {
	colToLook := col+magnitude
	return findChar(valToFind, row, colToLook)
}
