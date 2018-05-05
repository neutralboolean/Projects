//main.go

package main

import (
   "bufio"
   "errors"
   "fmt"
   "math"
   "os"
   "strconv"
)

var err error = errors.New("Variable not yet set.")
var numOfDigits int

func main () {
	scanner := bufio.NewScanner(os.Stdin)

	for (err != nil) {
		fmt.Println("How many digits (a positive integer) of 'e' are to be displayed?: ")
		scanner.Scan()
		numOfDigits, err = strconv.Atoi(scanner.Text())
		if (err == nil) {
			fmt.Printf("e out to %d decimal points is: %s\n", numOfDigits, strconv.FormatFloat(math.E, 'f', numOfDigits, 64)) 
		} else {
			fmt.Println("Input was not a valid integer.\n\t=========\n")
		}
	}
}