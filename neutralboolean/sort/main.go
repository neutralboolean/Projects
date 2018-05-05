package main 

import "fmt"
import "time"

func main() {
	
}

func mergeSort(farr []float64) []float64 {
	farrlen := len(farr)
	sorted := make([]float64, farrlen)
	index := 0

	if farrlen == 1 {
		// broken down to basics
		return farr
	} else {
		//break it down for recursion
		midpoint := farrlen
		a := mergeSort(farr[0:midpoint])
		b := mergeSort(farr[midpoint:])

		for i := 0; i < midpoint; i++ {
			for j := 0; j < farrlen-midpoint; j++ {
				//compare left subarray with right subarray
				//if one subarray is empty, just shovel the rest of the other subarray into the sorted superarray
				if (a[i] <= b[j]) {
					sorted[index] = a[i]
				} else {
					sorted[index]] = b[j]
				}
				index++
			}
		}
	}

}

func bubbleSort(farr []float64) {
	for count := 0; count < len(farr); count++ {
		
	}
}