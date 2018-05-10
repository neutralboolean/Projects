package main 

import (
    "fmt"
    "strconv"
)

func main() {
	var num, result int

	fmt.Print("Input the credit card number: ")
	fmt.Scanf("%d", &num)

	str := strconv.Itoa(num)

	check, err := strconv.Atoi(string(str[len(str)-1]))
	doubledigit := true 		//instead of doing calculations in the `for` headers,
								//I'm just going to use a flipping bit to check if I should double the current digit
	if err == nil {
		result = 0
		for i := len(str)-2; i >= 0; i-- {
			strfromrune := string(str[i])
			temp, err := strconv.Atoi(strfromrune)

			if doubledigit {
				temp *= 2
				if temp > 9 && err == nil {
					temp -= 9
				}

				doubledigit = false
			} else {
				doubledigit = true
			}

			// fmt.Print(temp, " ")

			result += temp
		}

		// fmt.Println("result = ", result)

		newstr := strconv.Itoa(result*9)

		result += check

		if result%10 == 0 {
			// fmt.Println(newstr, " ", str)

			if newstr[len(newstr)-1] == str[len(str)-1] {
				fmt.Printf("The credit card number given (%d) is valid.\n", num)
			} else {
				fmt.Printf("Failed the check digit test: The credit card number given (%d) is invalid.\n", num)
			}
		} else {
			fmt.Printf("Failed the modulo test: The credit card number given (%d) is invalid.\n", num)
		}
	}
}