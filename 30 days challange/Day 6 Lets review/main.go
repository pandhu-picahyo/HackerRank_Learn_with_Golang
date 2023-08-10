package main

import (
	"fmt"
	// start of new import
	"bufio"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	//Start of code

	reader := bufio.NewReader(os.Stdin)
	t, _ := reader.ReadString('\n')
	t = strings.TrimSpace(t)
	T := parseInt(t)

	for i := 0; i < T; i++ {
		S, _ := reader.ReadString('\n')
		S = strings.TrimSpace(S)

		var even, odd string
		for j, ch := range S {
			if j%2 == 0 {
				even += string(ch)
			} else {
				odd += string(ch)
			}
		}

		fmt.Printf("%s %s\n", even, odd)
	}
}

func parseInt(s string) int {
	var res int
	for _, ch := range s {
		res = res*10 + int(ch-'0')
	}
	return res
}

// end of code
