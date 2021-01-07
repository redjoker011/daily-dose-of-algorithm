// Package main provides ...
package main

import "fmt"

var set = []int{0, 1, 2}
var subset []int
var n = len(set)

func main() {
	for _, v := range set {
		search(v)
	}
}

// Generate Subset recursively
// https://stackoverflow.com/a/56925348/8680724
func search(k int) {
	if k == n {
		// Process Subsets
		if len(subset) >= 1 {
			fmt.Printf("Subset %v\n", subset)
		}
	} else {
		search(k + 1)
		subset = append(subset, k)
		search(k + 1)
		subset = subset[:len(subset)-1]
	}
}
