// Package main provides ...
package main

import (
	"fmt"
	"math"
	"sort"
)

var collection = []int{1, 2, 9, 2, 6}

// Minimizing Sum Algorithm
func main() {
	caseC1()
	caseC2()
}

// Case c = 1
// |a1 −x|+|a2 −x|+···+|an −x|.
// Get the median, iterate over collection and get the difference then get the
// sum of each iteration
func caseC1() {
	sort.Ints(collection)
	middlePointer := len(collection) / 2
	middleValue := collection[middlePointer]

	var sum int
	for _, v := range collection {
		sum += v - middleValue
	}

	fmt.Printf("Case C1 Sum %v\n", sum)
}

// Case c = 2
// (a1−x) +(a2−x) +···+(an−x) .
// Get the average value of collection and get the sum of the difference of
// average and n into power of 2
func caseC2() {
	collectionSum := 0

	for _, v := range collection {
		collectionSum += v
	}

	average := collectionSum / len(collection)

	var sum int
	for _, v := range collection {
		sum += int(math.Pow(float64(v-average), 2))
	}
	fmt.Printf("Case C2 Sum %v\n", sum)
}
