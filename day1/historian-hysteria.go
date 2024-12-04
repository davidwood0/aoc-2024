package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read in the input into two arrays
	arr1, arr2, err := readFile("input.txt");

	if err!=nil{
		os.Exit(1)
	}

	sortLeft := mergeSort(arr1);
	sortRight := mergeSort(arr2);

	outFile, err := os.Create("output.txt")
	if err != nil {
		os.Exit(1)
	}
	defer outFile.Close()

	result := 0;
	for i := 0; i < len(sortLeft); i++ {
		result += absDiff(sortLeft[i], sortRight[i])
		fmt.Fprintln(outFile, sortLeft[i], "  ", sortRight[i], "   diff: ", absDiff(sortLeft[i], sortRight[i]))
	}
	fmt.Fprintln(outFile, "total diff: ", result)

	return
}

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array;
	}

	half := len(array)/2;
	// split the array in half and merge sort each half
	firstHalf := mergeSort(array[:half])
	secondHalf := mergeSort(array[half:])

	return merge(firstHalf, secondHalf)
}

func merge(a []int, b []int) []int {
	result := make([]int, 0)

	i := 0;
	j := 0;

	for i < len(a) && j < len(b) {
		if (a[i] < b[j]) {
			// a's item is smaller than b's, append a to result
			result = append(result, a[i])
			i++;
		} else {
			// a's item is gtr than or equal to b, append b to result
			result = append(result, b[j])
			j++;
		}
	}

	for ; i < len(a); i++ {
		result = append(result, a[i])
	}

	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result;
}

func readFile(inFile string) ([]int, []int, error) {
    f, err := os.Open(inFile)
    if err != nil {
        fmt.Println(err)
        return nil,nil, err
    }
    
    r := bufio.NewReader(f)
    
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	for {
	    line, err := r.ReadString('\n')
		// EOF is treated as an error, but we still have last line of data left to process
		// so don't break on that
		// break on any other un-expected errors though
		if err != nil && err != io.EOF {
			break;
	    }
	
	    split := strings.Split(line, "   ")

		int1, _ := strconv.Atoi(strings.TrimSpace(split[0]))
		int2, _ := strconv.Atoi(strings.TrimSpace(split[1]))	

		arr1 = append(arr1, int1)
		arr2 = append(arr2, int2)

		// break if we hit EOF
		if err == io.EOF {
			break;
		}
	}
	
	defer f.Close()

	return arr1, arr2, nil
}

// the difference between a and b, absolute
func absDiff(a int, b int) int {
	if (a - b) < 0 {
		return (a-b)*-1;
	} else {
		return a-b;
	}

}