// Package main provides ...
package main

import "fmt"

func main() {
	var collection = []int{-1, 1, 2, 3, 5}
	const x = 5

	leftPointer := 0
	rightPointer := len(collection) - 1
	for leftPointer <= rightPointer {
		sum := collection[leftPointer] + collection[rightPointer]
		if sum < x {
			leftPointer = leftPointer + 1
		} else if sum > x {
			rightPointer = rightPointer - 1
		} else {
			fmt.Printf("Sum of %v is %v and %v", x, collection[leftPointer], collection[rightPointer])
			break
		}
	}
}
