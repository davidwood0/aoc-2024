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

	result,_ := calcMultiply("input.txt")

	println(result)
}

func calcMultiply(inFile string) (int, error) {
	f, err := os.Open(inFile)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	r := bufio.NewReader(f)

	result := 0;

	for {
		line, err := r.ReadString('\n')
		// EOF is treated as an error, but we still have last line of data left to process
		// so don't break on that
		// break on any other un-expected errors though
		if err != nil && err != io.EOF {
			break
		}

		split := strings.Split(line, "mul(")

		for _, v := range split {
			// then split by )
			split2 := strings.Split(v, ")")

			if (len(split2)>0) {
				// validate that int,int is all that is left, anything else, discard
				result += extractVal(split2[0])
			}
		}

		// break if we hit EOF
		if err == io.EOF {
			break
		}
	}

	defer f.Close()

	return result, nil;
}

func extractVal(line string) int {

	split := strings.Split(line, ",")

	if (len(split)==2) {
		int1, err := strconv.Atoi(split[0])
		if (err != nil) {
			return 0
		}

		int2, err := strconv.Atoi(split[1])
		if (err != nil) {
			return 0
		}

		if (int1 < 1000 && int2 < 1000 && int1 > -1 && int2 > -1) {
			// valid ints at both places, multiply
			return int1*int2;
		}

		return 0
	}
	return 0
}

