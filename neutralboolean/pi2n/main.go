package main 

import (
   "bufio"
   "errors"
   "fmt"
   "strconv"
   "math"
   "os"
)

var err error = errors.New("Variable was not set.")
var numOfDigits int

func main () {
	scanner := bufio.NewScanner(os.Stdin)

	for (err != nil) {
		fmt.Println("How many digits (a positive integer) of pi are to be displayed?: ")
		scanner.Scan()
		numOfDigits, err = strconv.Atoi(scanner.Text())
		if (err == nil) {
			fmt.Printf("pi out to %d decimal points is: %s\n", numOfDigits, strconv.FormatFloat(math.Pi, 'f', numOfDigits, 64)) 
		} else {
			fmt.Println("Input was not a valid integer.\n\t=========\n")
		}
	}
}
