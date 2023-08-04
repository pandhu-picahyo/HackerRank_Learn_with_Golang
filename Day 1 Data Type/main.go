package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var nilai1 uint64
	var nilai2 float64
	var nilai3 string

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	nilai1, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	nilai2, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	nilai3 = scanner.Text()
	// Print the sum of both integer variables on a new line.

	fmt.Println(i + nilai1)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+nilai2)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + nilai3)
}
