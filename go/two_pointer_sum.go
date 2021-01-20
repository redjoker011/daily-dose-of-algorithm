// Package main provides ...
package main

import "fmt"

func main() {
	easy()
	medium()
}

// Find two number that sum is equal to x
// Set two pointers
// left - lower index
// right - higher index
// Iterate over collection and move pointer accordingly
// - Increase left when sum of left  and right is lower than x
// - Decrease left when sum of left  and right is higher than x
// - Break loop when sum matches x
func easy() {
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
			fmt.Printf("Two Number Combination of %v is %v and %v \n", x, collection[leftPointer], collection[rightPointer])
			break
		}
	}
}

// Find the three numbers that sum is equal to x
// Setup nested loop
// Iterate Over collection in outer loop and get the target by getting the difference between
// iterated element and x.
// Then get the sum of left and right pointer in the nested loop while
// continously increasing leftPointer until it reach the last element then
// continue the iteration to the next element in the outer loop
func medium() {
	var collection = []int{1, 2, 4, 5, 12}
	const x = 19

	rightPointer := len(collection) - 1
	for i, v := range collection {
		target := x - v
		leftPointer := i + 1
		for j := leftPointer; j <= rightPointer; j++ {
			sum := collection[j] + collection[rightPointer]

			if sum == target {
				fmt.Printf("Three Number Combination of %v is %v, %v and %v \n", x, collection[j], collection[rightPointer], v)
			} else {
				leftPointer++
			}
		}
	}
}
