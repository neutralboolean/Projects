package main 

import "fmt"

func main() {
	var input string
	fmt.Println("Input string to pig-latinize: ")
	fmt.Scanln(&input)
	piggy := fmt.Sprintf("%s-%cay", input[1:], input[0])
	fmt.Printf("%q in pig latin is %q\n", input, piggy)
}