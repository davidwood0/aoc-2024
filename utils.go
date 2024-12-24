package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadFileIntoArrayOfIntArray(inFile string) ([][]int, error) {
	f, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r := bufio.NewReader(f)

	arr1 := make([]int, 0)
	arr2 := make([][]int, 0)

	for {
		line, err := r.ReadString('\n')
		// EOF is treated as an error, but we still have last line of data left to process
		// so don't break on that
		// break on any other un-expected errors though
		if err != nil && err != io.EOF {
			break
		}

		split := strings.Split(line, " ")

		for _, v := range split {
			int1, _ := strconv.Atoi(strings.TrimSpace(v))
			arr1 = append(arr1, int1)
		}

		arr2 = append(arr2, arr1)

		arr1 = make([]int, 0)

		// break if we hit EOF
		if err == io.EOF {
			break
		}
	}

	defer f.Close()

	return arr2, nil
}




