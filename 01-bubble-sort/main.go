package main

import (
	"fmt"
)

// BubbleSort - takes in an input which is a slice of ints and returns the sorted slice
func BubbleSort(input []int) []int {
	// n is the number of items in our list
	n := len(input)
	// set swapped to true
	swapped := true
	// loop
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 0; i < n-1; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i] > input[i+1] {
				// log that we are swapping values for posterity
				fmt.Println("Swapping")
				// swap values using Go's tuple assignment
				input[i], input[i+1] = input[i+1], input[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}

	return input
}

func main() {
	fmt.Println("Bubble Sorting Algorithm in Go")

	unsortedInput := []int{5, 3, 4, 1, 2}
	sorted := BubbleSort(unsortedInput)

	fmt.Println(sorted)
}
