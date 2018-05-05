package main 

import "fmt"

func main() {
	var input string

	fmt.Println("Input some string to reverse. End by pressing the Enter key: ")
	fmt.Scanln(&input)

	reversed := make([]rune, len(input))
	lindex := len(input) - 1
	for i, v := range(input) {
		reversed[lindex-i] = v
	}

	fmt.Printf("\"%s\" reversed is \"%s\"\n", input, string(reversed))
}