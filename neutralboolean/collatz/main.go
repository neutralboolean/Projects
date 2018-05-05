package main 

import "fmt"

func main() {	
	var input int

	fmt.Println("What number would you like to start from?: ")
	fmt.Scanf("%d\n", &input)
	start := input 
	count := 0

	for input > 1 {
		if (input % 2) == 0 {
			fmt.Print("Even:\t")
			fmt.Print(input)
			input /= 2
			fmt.Printf(" =>> %d\n", input)
		} else {
			fmt.Print("Odd:\t")
			fmt.Print(input)
			input = (input * 3) + 1
			fmt.Printf(" =>> %d\n", input)
		}
		count++
	}

	fmt.Printf("It takes %d steps to get from %d to 1\n", count, start)
}