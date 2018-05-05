package main 

import (
   "bufio"
   "errors"
   "fmt"
   "os"
   "strconv"
)

/*
Possible improvements:
- dynamic programming/recording known primes for any number to speed up prcoessing
- sort the results, just for beautification
*/

func main() {
	primeFactors := make([]int, 0) 
	scanner := bufio.NewScanner(os.Stdin)

	if err := errors.New("Not yet initialized."); err != nil {
		fmt.Println("Find the prime factor of which number?: ")
		
		if scanner.Scan() {
			input, err := strconv.Atoi(scanner.Text())
			if err == nil { 
				if Factorize(input, &primeFactors) == nil {
					printFactors(input, primeFactors)
				}
			}
		}
	}
}

//recursively determines list of prime factors
//inputs: `num`, the num whose factors are to be returned; `lof`, the 'list of factors' that is actually returned
func Factorize(num int, lof *[]int) error {
	if num < 1 {
		return errors.New("Tried to find the prime factor of a nonpositive number.")
	}

	isPrime := true

	for i := num-1; i > 1; i-- {
		if (num % i) == 0 {
			f0 := i
			f1 := num / i
			// find the prime factors of these two non-prime factors
			Factorize(f0, lof)
			Factorize(f1, lof)
			isPrime = false
			break
		}
	}

	if isPrime { *lof = append(*lof, num) }
	return nil 
}

func printFactors(num int, intarr []int) {
	fmt.Printf("\nThe factors of %d are: ", num)

	for i, v := range(intarr) {
		if i == len(intarr)-1 {
			fmt.Println(v)
		} else {
			fmt.Printf("%d, ", v)
		}
	}
}