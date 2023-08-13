package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// Start of the code
	reader := bufio.NewReader(os.Stdin)

	// Read the number of entries in the phone book
	nTemp, _ := reader.ReadString('\n')
	nTemp = strings.TrimSpace(nTemp)
	n := parseInt(nTemp)

	// Create a phone book (map) to store names and phone numbers
	phoneBook := make(map[string]string)

	// Read and store phone book entries
	for i := 0; i < n; i++ {
		entry := strings.Fields(readLine(reader))
		name := entry[0]
		phoneNumber := entry[1]
		phoneBook[name] = phoneNumber
	}

	// Read and process queries
	for {
		query, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		query = strings.TrimSpace(query)
		if phoneNumber, ok := phoneBook[query]; ok {
			fmt.Printf("%s=%s\n", query, phoneNumber)
		} else {
			fmt.Println("Not found")
		}
		if err == io.EOF {
			break
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return string(str)
}

func parseInt(s string) int {
	var res int
	for _, ch := range s {
		res = res*10 + int(ch-'0')
	}
	return res

	// End of code
}
