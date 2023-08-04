package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)

	var datainput string

	scanner.Scan()
	datainput = scanner.Text()
	fmt.Println("Hello, World.")
	fmt.Println(datainput)
}
