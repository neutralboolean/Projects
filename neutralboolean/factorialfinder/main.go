package main 

import "fmt"

//only calculates up to `65!`

func main() {
	var n uint64
	fmt.Println("Input non-negative integer for `n` in factorial `n!`.")
	fmt.Scanf("%d", &n)

	fmt.Printf("Factorial calculated via loop:\n%d! = %d\n", n, LoopFactorial(n))
	fmt.Printf("Factorial calculated via recursion:\n%d! = %d\n", n, RecursiveFactorial(n))
}

func LoopFactorial(n uint64) uint64 {
	var result uint64 = 1
	var i uint64

	for i = 1;i <= n; i++ {
		result *= i
	}

	return result
}

func RecursiveFactorial(n uint64) uint64 {
	if (n > 1) {
		return n * RecursiveFactorial(n-1)
	}

	return 1
}