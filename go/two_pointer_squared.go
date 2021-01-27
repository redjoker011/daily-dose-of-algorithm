// Package main provides ...
package main

import (
	"fmt"
	"math"
)

func main() {
	var collection = []int{-4, -3, 1, 2, 3}

	rightPointer := 0

	// Set the right pointer to the minimum positive value
	// Iterate over collection until value is greater than 0
	for rightPointer < len(collection) && collection[rightPointer] < 0 {
		rightPointer++
	}

	fmt.Printf("R Pointer Set to %v\n", rightPointer)
	leftPointer := rightPointer - 1
	fmt.Printf("L Pointer Set to %v\n", leftPointer)

	var output []int

	// Iterate Over as long as left is not out of bounce and right is not out of
	// bounce of array
	for leftPointer >= 0 && rightPointer < len(collection) {
		leftValue := int(math.Pow(float64(collection[leftPointer]), 2))
		rightValue := int(math.Pow(float64(collection[rightPointer]), 2))

		// Compare squared values and append minimum value into the output
		// collection
		// When left value is lower append value into collection and decrement
		// pointer
		// When right value is lower append value into collection and increment
		// pointer
		if leftValue < rightValue {
			output = append(output, leftValue)
			leftPointer--
		} else {
			output = append(output, rightValue)
			rightPointer++
		}
	}

	// Check Left and Right pointer to ensure all elements were included in the
	// collection
	for leftPointer >= 0 {
		output = append(output, int(math.Pow(float64(collection[leftPointer]), 2)))
		leftPointer--
	}
	for rightPointer < len(collection) {
		output = append(output, int(math.Pow(float64(collection[rightPointer]), 2)))
		rightPointer++
	}

	fmt.Printf("Sorted Squared Values %v\n", output)
}
