package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var rowMap = make(map[string]map[int]map[int]bool)

var oppositeMap map[string]string = map[string]string{
	"M": "S",
	"S": "M",
}

type searchDir func(valToFind string, row int, col int) bool

var funcOppositeMap map[string]searchDir = map[string]searchDir{
	"NW": findSouthEast,
	"NE": findSouthWest,
}

func main() {

	err := mapRows("input.txt")

	if err != nil {
		os.Exit(1)
	}

	result := solveForLetter("A")

	println(result)

}

func solveForLetter(letter string) int {

	result := 0;
	resultMap := map[string]int{};

	letterMap := rowMap[letter];

	for row, colMap := range letterMap {
		for col := range colMap {

			 if (checkCorners(row, col)){
				result++;
			 }
			
		}
	}

	for i,v := range resultMap {
		println(fmt.Sprintf("%v results: %v", i, v))
	}

	return result;
}

func checkCorners(row int, col int) bool {
	return checkNeDiagCorners(row, col) && checkNwDiagCorners(row, col)
}

func checkNwDiagCorners(row int, col int) bool {

	valFound := ""

	if (findNorthWest("M", row, col)){
		valFound = "M"
	} else if (findNorthWest("S", row, col)) {
		valFound = "S"
	} else {
		return false;
	}

	return funcOppositeMap["NW"](oppositeMap[valFound], row, col)
}

func checkNeDiagCorners(row int, col int) bool {

	valFound := ""

	if (findNorthEast("M", row, col)){
		valFound = "M"
	} else if (findNorthEast("S", row, col)) {
		valFound = "S"
	} else {
		return false;
	}

	return funcOppositeMap["NE"](oppositeMap[valFound], row, col)
}

func findChar(valToFind string, rowToLook int, colToLook int) bool {
	found := rowMap[valToFind][rowToLook][colToLook];

	return found;
}

func findNorthWest(valToFind string, row int, col int) bool {
	return findChar(valToFind, row-1, col-1)
}

func findNorthEast(valToFind string, row int, col int) bool {
	return findChar(valToFind, row-1, col+1)
}

func findSouthWest(valToFind string, row int, col int) bool {
	return findChar(valToFind, row+1, col-1)
}

func findSouthEast(valToFind string, row int, col int) bool {
	return findChar(valToFind, row+1, col+1)
}

func mapRows(inFile string) error {
	f, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	r := bufio.NewReader(f)

	rowCount := 0

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

			validLetter := false

			switch v {
			case "M", "A", "S": validLetter = true
			}

			if validLetter {
				val, ok := letterColMap[v]
				if !ok {
					letterColMap[v] = map[int]bool{colCount: true}
				} else {
					val[colCount] = true
				}
			}
			colCount++
		}

		for i, v := range letterColMap {
			val, ok := rowMap[i]
			if !ok {
				rowMap[i] = map[int]map[int]bool{rowCount: v}
			} else {
				val[rowCount] = v
			}
		}

		rowCount++

		// break if we hit EOF
		if err == io.EOF {
			break
		}
	}

	defer f.Close()

	return nil
}
