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

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	// start of code

	maxHourglassSum := calculateMaxHourglassSum(arr)
	fmt.Println((maxHourglassSum))
}

func calculateHourglassSum(arr [][]int32, row, col int) int32 {
	sum := arr[row][col] + arr[row][col+1] + arr[row][col+2] +
		arr[row+1][col+1] +
		arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]
	return sum
}

func calculateMaxHourglassSum(arr [][]int32) int32 {
	maxSum := int32(-63) // Initialize with the lowest possible value for hourglass sum

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := calculateHourglassSum(arr, row, col)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
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
