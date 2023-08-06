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

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	// write your condition here
	if isWeird(N) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
	// end of condition
}

// write your function here
func isWeird(n int32) bool {
	if n%2 != 0 {
		return true
	} else {
		if n >= 2 && n <= 5 {
			return false
		} else if n >= 6 && n <= 20 {
			return true
		} else {
			return false
		}
	}
}

// end of function

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
