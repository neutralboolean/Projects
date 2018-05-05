package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var err error = errors.New("Variable not yet set.")
var numOfDigits int
var f0, f1 int = 0, 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for err != nil {
		fmt.Println("Input integer for length of Fibonacci sequence: ")
		scanner.Scan()

		numOfDigits, err = strconv.Atoi(scanner.Text())
		if err == nil {
			fmt.Printf("%d %d ", f0, f1)
			for i := 0; i < numOfDigits-2; i++ {
				temp := f1
				f1 = f0 + f1
				f0 = temp
				fmt.Printf("%d ", f1)
			}
			fmt.Println()
		} else {
			fmt.Println("Input was not valid.\n\t=======\n")
		}
	}
}
