package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(array []int) []int {
	// if we have an array of size 1
	// we can just instantly return
	if len(array) < 2 {
		return array
	}

	// we want the left-most and the right-most index of
	// the array we are going to sort
	left, right := 0, len(array)-1

	// choose a random pivot
	pivotIndex := rand.Int() % len(array)

	//
	array[pivotIndex], array[right] = array[right], array[pivotIndex]

	// we iterate through every element of the array and
	// place the items smaller than the pivot on the left
	// and then the items greater than the pivot on the right
	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	// swap the left
	array[left], array[right] = array[right], array[left]

	// recursively call our quicksort algorithm on
	// the two distinct partitions
	QuickSort(array[:left])
	QuickSort(array[left+1:])

	// et voila, we return our sorted array
	return array
}

func main() {
	test := []int{5, 3, 2, 57, 34, 2314, 3, 2, 1}
	sorted := QuickSort(test)
	fmt.Println(sorted)
}
