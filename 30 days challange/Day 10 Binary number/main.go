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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	// start of code
	// Convert the number to binary and find the maximum consecutive 1's
	binaryStr := strconv.FormatInt(int64(n), 2)
	maxConsecutive := findMaxConsecutiveOnes(binaryStr)

	fmt.Println(maxConsecutive)

}

func findMaxConsecutiveOnes(binaryStr string) int {
	maxConsecutive := 0
	currentConsecutive := 0

	for _, digit := range binaryStr {
		if digit == '1' {
			currentConsecutive++
			if currentConsecutive > maxConsecutive {
				maxConsecutive = currentConsecutive
			}
		} else {
			currentConsecutive = 0
		}
	}

	return maxConsecutive
}

// end of code

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
